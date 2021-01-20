package taskgroup

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"syscall"
)

const (
	STATUS_Running    = "running"
	STATUS_Cancelled  = "cancelled"
	STATUS_Done       = "done"
	STATUS_Failed     = "failed"
	STATUS_Unexpected = "unexpected"

	STAGE_Starting  = "starting"  // created and add task
	STAGE_Running   = "running"   // waiting
	STAGE_Canceling = "canceling" // call taskgroup.Cancel
	STAGE_Exit      = "exit"      // one of tasks status is cancelled or failed
	STAGE_Done      = "done"      // all tasks are done
)

type Task struct {
	idx    int
	Name   string `json:"name"`            // task name
	Status string `json:"status"`          // task status: ["running","cancelled","done","failed"]
	Error  error  `json:"error,omitempty"` // task error for status failed/cancelled/unexpected
	// Start  *time.Time `json:"start,omitempty"` // start time
	// End    *time.TIme `json:"end,omitempty"`   // end time
}

type TaskGroup struct {
	name        string
	debug       bool
	tks         []Task
	ctx         context.Context
	cancel      func()
	once, once2 *sync.Once
	mutex       *sync.Mutex
	ch          chan Task
	stage       string // starting, running (waiting), canceling, exit/done/unexpected
}

func GetFuncName(v interface{}) (pkg, name string) {
	str := runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name()
	return filepath.Dir(str), filepath.Base(str)
}

////
func (tk *Task) GetIndex() int {
	return tk.idx
}

func (tk *Task) String() string {
	msg := "nil"
	if tk.Error != nil {
		msg = tk.Error.Error()
	}

	return fmt.Sprintf(
		"index: %d, name: %q, status: %q, error: %q",
		tk.idx, tk.Name, tk.Status, msg,
	)
}

func (tk *Task) PrintError() {
	if tk.Error == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "%s\n", tk.String())
}

////
// create *TaskGroup
func NewTaskGroup(c context.Context, name string, debug bool) (tg *TaskGroup) {
	tg = new(TaskGroup)

	tg.debug = debug
	tg.tks = make([]Task, 0, 3)
	tg.ctx, tg.cancel = context.WithCancel(c)
	tg.once, tg.once2 = new(sync.Once), new(sync.Once)
	tg.mutex = new(sync.Mutex)
	tg.ch = make(chan Task)
	tg.stage = STAGE_Starting
	return
}

func (tg *TaskGroup) GetName() string {
	return tg.name
}

func (tg *TaskGroup) addTask(name string) (tk Task) {
	tg.mutex.Lock()
	tk = Task{idx: len(tg.tks), Name: name, Status: STATUS_Running}
	if tk.Name == "" {
		tk.Name = fmt.Sprintf("Task[%d]", tk.idx)
	}
	tg.tks = append(tg.tks, tk)
	tg.mutex.Unlock()

	tk.Status = STATUS_Done // default value
	return tk
}

func (tg *TaskGroup) log(format string, a ...interface{}) {
	if !tg.debug {
		return
	}
	fmt.Printf(">>> "+strings.TrimSpace(format)+"\n", a...)
}

// run task
func (tg *TaskGroup) Go(run func() error, stop func(), names ...string) {
	var tk Task

	if len(names) > 0 {
		tk = tg.addTask(names[0])
	} else {
		_, funcName := GetFuncName(run)
		tk = tg.addTask(funcName)
	}

	chErr := make(chan error)
	tg.log("TaskGroup.Go[%q] step A: starting", tk.Name)

	go func() {
		var err error

		defer func() {
			if v := recover(); v != nil {
				err = fmt.Errorf("panic: %v", v)
				tk.Status = STATUS_Unexpected
			}

			tg.log("TaskGroup.Go[%q] step B: return %v", tk.Name, err)
			chErr <- err
		}()

		if err = run(); err != nil && tk.Status != STATUS_Cancelled {
			tk.Status = STATUS_Failed
		}
	}() // goruntime1

	go func() {
		<-tg.ctx.Done()
		tg.log("TaskGroup.Go[%q] step C: received taskgroup cancelled", tk.Name)
		tk.Status = STATUS_Cancelled

		if stop != nil {
			tg.log("TaskGroup.Go[%q] step D: execute task cancel()", tk.Name)
			stop() // must cause task() return
		} else {
			tg.log("TaskGroup.Go[%q] step E: send cancel error to channel", tk.Name)
			chErr <- fmt.Errorf("canncelled by taskgroup") //!!! may cause goruntime1 leak
		}
	}()

	go func() {
		tk.Error = <-chErr
		tg.log(
			"TaskGroup.Go[%q] step F: send task(%q, %v) result to taskgroup",
			tk.Name, tk.Status, tk.Error,
		)

		tg.ch <- tk
	}()
}

// run a function with context
func (tg *TaskGroup) GoWithContext(run func(context.Context) error, names ...string) {
	var tk Task

	if len(names) > 0 {
		tk = tg.addTask(names[0])
	} else {
		_, funcName := GetFuncName(run)
		tk = tg.addTask(funcName)
	}
	tg.log("TaskGroup.GoWithContext[%q] step A: starting", tk.Name)

	go func() {
		defer func() {
			if v := recover(); v != nil {
				tk.Error, tk.Status = fmt.Errorf("panic: %v", v), STATUS_Unexpected
			}

			tg.log(
				"TaskGroup.GoWithContext[%q] step B: task(%q, %v) return",
				tk.Name, tk.Status, tk.Error,
			)
			tg.ch <- tk // no status "cancelled"
		}()

		if tk.Error = run(tg.ctx); tk.Error != nil {
			tk.Status = STATUS_Failed
		}
	}()
}

// get number of tasks
func (tg *TaskGroup) NumOfTasks() int {
	return len(tg.tks)
}

func (tg *TaskGroup) GetStage() string {
	return tg.stage
}

// copy tasks
func (tg *TaskGroup) GetTasks() (tks []Task) {
	tks = make([]Task, len(tg.tks))
	copy(tks, tg.tks)
	return tks
}

// wait all tasks exit
func (tg *TaskGroup) wait() {
	tg.stage = STAGE_Running

	for i := 0; i < len(tg.tks); i++ {
		tk := <-tg.ch
		tg.tks[tk.idx] = tk
	}

	tg.stage = STAGE_Done
	for i := range tg.tks {
		if tg.tks[i].Status != STATUS_Done {
			tg.stage = STAGE_Exit
			break
		}
	}
}

func (tg *TaskGroup) Wait() {
	tg.once2.Do(func() { tg.wait() })
}

// cancel all tasks
func (tg *TaskGroup) Cancel() {
	tg.once.Do(func() {
		tg.stage = STAGE_Canceling
		tg.cancel()
	})

	tg.Wait()
}

// listen os interrupt signals, and cancel all task
func (tg *TaskGroup) ListenOSIntr(sgs ...os.Signal) (err error) {
	chErr := make(chan error)

	go func() {
		tg.Wait()
		chErr <- nil
	}()

	go func() {
		quit := make(chan os.Signal)

		if len(sgs) == 0 {
			signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
		} else {
			signal.Notify(quit, sgs...)
		}

		<-quit
		tg.Cancel()
		chErr <- fmt.Errorf("tasks was intrrupted")
	}()

	return <-chErr
}
