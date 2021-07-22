package misc

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Excutor struct {
	exitFuncs []func(error)
	errors    []error
	lock      *sync.Mutex
	once      *sync.Once
	ch        chan bool
}

func NewExcutor() *Excutor {
	return &Excutor{
		exitFuncs: make([]func(error), 0),
		errors:    make([]error, 0),
		lock:      new(sync.Mutex),
		once:      new(sync.Once),
		ch:        make(chan bool),
	}
}

func (ex *Excutor) setErr(idx int, err error) bool {
	ex.lock.Lock()
	defer ex.lock.Unlock()

	if idx >= len(ex.errors) {
		return false
	}
	ex.errors[idx] = err
	return true
}

func (ex *Excutor) getErr(idx int) (err error) {
	ex.lock.Lock()
	defer ex.lock.Unlock()

	if idx >= len(ex.errors) {
		return nil
	}
	return ex.errors[idx]
}

func (ex *Excutor) Load(run func() error, onExit func(error)) {
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

func (ex *Excutor) Wait(dura time.Duration, sgs ...os.Signal) {
	quit := make(chan os.Signal)

	if len(sgs) == 0 {
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	} else {
		signal.Notify(quit, sgs...)
	}

	select {
	case sig := <-quit:
		log.Printf("Excutor: received signal: %v\n", sig)
	case <-ex.ch:
		log.Printf("Excutor: task(s) failed")
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
