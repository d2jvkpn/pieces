package misc

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Executor struct {
	exitFuncs []func(error)
	errors    []error
	lock      *sync.Mutex
	once      *sync.Once
	ch        chan bool
}

func NewExecutor() *Executor {
	return &Executor{
		exitFuncs: make([]func(error), 0),
		errors:    make([]error, 0),
		lock:      new(sync.Mutex),
		once:      new(sync.Once),
		ch:        make(chan bool),
	}
}

func (ex *Executor) setErr(idx int, err error) bool {
	ex.lock.Lock()
	defer ex.lock.Unlock()

	if idx >= len(ex.errors) {
		return false
	}
	ex.errors[idx] = err
	return true
}

func (ex *Executor) getErr(idx int) (err error) {
	ex.lock.Lock()
	defer ex.lock.Unlock()

	if idx >= len(ex.errors) {
		return nil
	}
	return ex.errors[idx]
}

func (ex *Executor) Load(run func() error, onExit func(error)) {
	ex.exitFuncs = append(ex.exitFuncs, onExit)
	ex.errors = append(ex.errors, nil)
	index := len(ex.exitFuncs) - 1

	go func() {
		err := run()
		ex.setErr(index, err)
		ex.once.Do(func() {
			ex.ch <- false
		})
	}()
}

func (ex *Executor) Wait(dura time.Duration, sgs ...os.Signal) {
	quit := make(chan os.Signal)

	if len(sgs) == 0 {
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	} else {
		signal.Notify(quit, sgs...)
	}

	select {
	case sig := <-quit:
		log.Printf("Executor: received signal: %v\n", sig)
	case <-ex.ch:
		log.Printf("Executor: task(s) failed")
		if dura > 0 {
			time.Sleep(dura)
		}
	}

	for i := range ex.exitFuncs {
		if ex.exitFuncs[i] == nil {
			continue
		}
		ex.exitFuncs[i](ex.getErr(i))
	}
}
