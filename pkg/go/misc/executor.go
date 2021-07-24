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
		if err := run(); err != nil {
			log.Printf("Executor error %d: %v\n", index, err)
			ex.setErr(index, err)
		}

		ex.once.Do(func() {
			ex.ch <- false
		})
	}()
}

func (ex *Executor) Wait(dura time.Duration, sgs ...os.Signal) (ok bool) {
	var err error
	ok = true
	quit := make(chan os.Signal)

	if len(sgs) == 0 {
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	} else {
		signal.Notify(quit, sgs...)
	}

	select {
	case <-quit:
		log.Printf("Executor exit")
		ok = false
	case <-ex.ch:
		log.Println("Executor failed")
		if dura > 0 {
			time.Sleep(dura)
		}
	}

	for i := range ex.exitFuncs {
		if err = ex.getErr(i); err != nil {
			ok = false
		}
		if ex.exitFuncs[i] != nil {
			ex.exitFuncs[i](err)
		}
	}

	return ok
}
