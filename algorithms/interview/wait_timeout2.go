package interview

import (
	"fmt"
	"sync"
	"time"
)

type TimeoutWaitGroup struct {
	sync.WaitGroup
	ch   chan bool
	dura time.Duration
}

func NewTimeoutWaitGroup(dura time.Duration) (wg *TimeoutWaitGroup) {
	return &TimeoutWaitGroup{
		ch:   make(chan bool),
		dura: dura,
	}
}

func (wg *TimeoutWaitGroup) WaitTimeout() bool {
	go func() {
		wg.Wait()
		wg.ch <- true
	}()

	select {
	case <-wg.ch:
		return true
	case <-time.After(wg.dura):
		return false
	}
}

func (wg *TimeoutWaitGroup) AddJob(job func()) {
	wg.Add(1)
	go func() {
		fmt.Println("    >>>")
		job()
		fmt.Println("    <<<")
		wg.Done()
	}()
}

func Waittimeout2() {
	fmt.Println(">>> Waittimeout2:")
	wg := NewTimeoutWaitGroup(20 * time.Second)

	for i := 0; i < 10; i++ {
		wg.AddJob(job2)
	}

	if ok := wg.WaitTimeout(); ok {
		fmt.Println("    all jobs done.")
	} else {
		fmt.Println("    timeout!")
	}
}

func job2() {
	time.Sleep(5 * time.Second)
}
