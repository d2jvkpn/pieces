package main

import (
	"fmt"
	"time"
)

type FlowLimiter struct {
	dur    time.Duration
	num    int64
	ch     chan bool
	ticker *time.Ticker
}

func rfc999ms() string {
	return time.Now().Format("2006-01-02T15:04:05.000Z07:00")
}

func NewFlowLimiter(dur time.Duration, num int64) (fl *FlowLimiter, err error) {
	if dur <= 0 || num <= 0 {
		return nil, fmt.Errorf("invalid dur or num")
	}

	fl = &FlowLimiter{
		dur:    dur,
		num:    num,
		ch:     make(chan bool, num),
		ticker: time.NewTicker(dur),
	}

	go func() {
		defer func() {
			if v := recover(); v != nil {
				fmt.Println("!!!", v) // send on closed channel
			}
		}()

		for _ = range fl.ticker.C {
			nl := len(fl.ch)
			fmt.Printf(">>> tick start: %s, len(chan): %d\n", rfc999ms(), nl)
			for i := nl; i < cap(fl.ch); i++ {
				fl.ch <- true // panic if fl.ch was closed
			}
			fmt.Printf("    tick end: %s, len(chan): %d\n", rfc999ms(), nl)
		}

		nl := len(fl.ch)
		fmt.Printf("=== tick stop: %s, len(chan): %d\n", rfc999ms(), nl)
	}()

	return
}

func (fl *FlowLimiter) String() string {
	return fmt.Sprintf("duration: %s, number: %d", fl.dur, fl.num)
}

func (fl *FlowLimiter) Get() bool {
	return <-fl.ch
}

func (fl *FlowLimiter) GetWithTimeout(timeout time.Duration) func(do func(), abort func()) {
	return func(do, abort func()) {
		ok := false

		select {
		case ok = <-fl.ch:
		case <-time.After(timeout):
		}

		if ok {
			do()
		} else if !ok && abort != nil {
			abort()
		}
	}
}

func (fl *FlowLimiter) Len() int {
	return len(fl.ch)
}

func (fl *FlowLimiter) ResetDuration(dur time.Duration) (err error) {
	if dur <= 0 {
		return fmt.Errorf("invalid dur")
	}

	fl.dur = dur
	fl.ticker.Reset(fl.dur)
	return
}

func (fl *FlowLimiter) Close() {
	fl.ticker.Stop()
	close(fl.ch)

	for i := 0; i < len(fl.ch); i++ { // clear exisiting values in chan
		<-fl.ch
	}
}
