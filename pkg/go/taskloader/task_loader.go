package TaskLoader

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
	"time"
)

const (
	STATUS_Running    = "running"
	STATUS_Cancelled  = "cancelled"
	STATUS_Done       = "done"
	STATUS_Failed     = "failed"
	STATUS_Unexpected = "unexpected"

	STAGE_Starting  = "starting"  // created and add task
	STAGE_Running   = "running"   // waiting
	STAGE_Canceling = "canceling" // call TaskLoader.Cancel
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

type TaskLoader struct {
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
// create *TaskLoader
func NewTaskLoader(c context.Context, name string, debug bool) (tl *TaskLoader) {
	tl = new(TaskLoader)

	tl.debug = debug
	tl.tks = make([]Task, 0, 3)
	tl.ctx, tl.cancel = context.WithCancel(c)
	tl.once, tl.once2 = new(sync.Once), new(sync.Once)
	tl.mutex = new(sync.Mutex)
	tl.ch = make(chan Task)
	tl.stage = STAGE_Starting
	return
}

func (tl *TaskLoader) GetName() string {
	return tl.name
}

func (tl *TaskLoader) addTask(name string) (tk Task) {
	tl.mutex.Lock()
	tk = Task{idx: len(tl.tks), Name: name, Status: STATUS_Running}
	if tk.Name == "" {
		tk.Name = fmt.Sprintf("Task[%d]", tk.idx)
	}
	tl.tks = append(tl.tks, tk)
	tl.mutex.Unlock()

	tk.Status = STATUS_Done // default value
	return tk
}

func (tl *TaskLoader) log(format string, a ...interface{}) {
	if !tl.debug {
		return
	}

	tmpl := fmt.Sprintf(
		">>> %s %s\n",
		time.Now().Format("2006-01-02T15:04:05.000Z07:00"),
		strings.TrimSpace(format),
	)

	fmt.Printf(tmpl, a...)
}

// run task
func (tl *TaskLoader) Go(run func() error, stop func(), names ...string) {
	var tk Task

	if len(names) > 0 {
		tk = tl.addTask(names[0])
	} else {
		_, funcName := GetFuncName(run)
		tk = tl.addTask(funcName)
	}

	chErr := make(chan error)
	tl.log("TaskLoader.Go[%q] step A: starting", tk.Name)

	go func() {
		var err error

		defer func() {
			if v := recover(); v != nil {
				err = fmt.Errorf("panic: %v", v)
				tk.Status = STATUS_Unexpected
			}

			tl.log("TaskLoader.Go[%q] step B: return %v", tk.Name, err)
			chErr <- err
		}()

		if err = run(); err != nil && tk.Status != STATUS_Cancelled {
			tk.Status = STATUS_Failed
		}
	}() // goruntime1

	go func() {
		<-tl.ctx.Done()
		tl.log("TaskLoader.Go[%q] step C: received TaskLoader cancelled", tk.Name)
		tk.Status = STATUS_Cancelled

		if stop != nil {
			tl.log("TaskLoader.Go[%q] step D: execute task cancel()", tk.Name)
			stop() // must cause task() return
		} else {
			tl.log("TaskLoader.Go[%q] step E: send cancel error to channel", tk.Name)
			chErr <- fmt.Errorf("canncelled by TaskLoader") //!!! may cause goruntime1 leak
		}
	}()

	go func() {
		tk.Error = <-chErr
		tl.log(
			"TaskLoader.Go[%q] step F: send task(%q, %v) result to TaskLoader",
			tk.Name, tk.Status, tk.Error,
		)

		tl.ch <- tk
	}()
}

// run a function with context
func (tl *TaskLoader) GoWithContext(run func(context.Context) error, names ...string) {
	var tk Task

	if len(names) > 0 {
		tk = tl.addTask(names[0])
	} else {
		_, funcName := GetFuncName(run)
		tk = tl.addTask(funcName)
	}
	tl.log("TaskLoader.GoWithContext[%q] step A: starting", tk.Name)

	go func() {
		defer func() {
			if v := recover(); v != nil {
				tk.Error, tk.Status = fmt.Errorf("panic: %v", v), STATUS_Unexpected
			}

			tl.log(
				"TaskLoader.GoWithContext[%q] step B: task(%q, %v) return",
				tk.Name, tk.Status, tk.Error,
			)
			tl.ch <- tk // no status "cancelled"
		}()

		if tk.Error = run(tl.ctx); tk.Error != nil {
			tk.Status = STATUS_Failed
		}
	}()
}

// get number of tasks
func (tl *TaskLoader) NumOfTasks() int {
	return len(tl.tks)
}

func (tl *TaskLoader) GetStage() string {
	return tl.stage
}

// copy tasks
func (tl *TaskLoader) GetTasks() (tks []Task) {
	tks = make([]Task, len(tl.tks))
	copy(tks, tl.tks)
	return tks
}

// wait all tasks exit
func (tl *TaskLoader) wait() {
	tl.stage = STAGE_Running

	for i := 0; i < len(tl.tks); i++ {
		tk := <-tl.ch
		tl.tks[tk.idx] = tk
	}

	tl.stage = STAGE_Done
	for i := range tl.tks {
		if tl.tks[i].Status != STATUS_Done {
			tl.stage = STAGE_Exit
			break
		}
	}
}

func (tl *TaskLoader) Wait() {
	tl.once2.Do(func() { tl.wait() })
}

// cancel all tasks
func (tl *TaskLoader) Cancel() {
	tl.once.Do(func() {
		tl.stage = STAGE_Canceling
		tl.cancel()
	})

	tl.Wait()
}

// listen os interrupt signals, and cancel all task
func (tl *TaskLoader) ListenOSIntr(sgs ...os.Signal) (err error) {
	chErr := make(chan error)

	go func() {
		tl.Wait()
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
		tl.Cancel()
		chErr <- fmt.Errorf("tasks was intrrupted")
	}()

	return <-chErr
}
