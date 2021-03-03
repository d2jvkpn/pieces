package main

import (
	"fmt"
	"sync"
)

// parallel counter with buffered and RWMutex
type Counter struct {
	m      *sync.RWMutex
	ch     chan int
	status bool
	n, s   int
}

func NewCounter(n int) (counter *Counter, err error) {
	if n < 1 {
		err = fmt.Errorf("invalid number for buffered channel")
		return
	}

	counter = &Counter{
		new(sync.RWMutex), make(chan int, n), true, n, 0,
	}

	for i := 0; i < n; i++ {
		counter.ch <- 0
	}

	return
}

// add a number
func (counter *Counter) Add(n int) (err error) {
	counter.m.RLock()
	defer counter.m.RUnlock()

	if !counter.status {
		// println("channel closed for", n)
		err = fmt.Errorf("channel closed")
		return
	}

	counter.ch <- (<-counter.ch) + n

	return
}

// get channel length and capcity
func (counter *Counter) GetN() (l, c int) {
	return len(counter.ch), counter.n
}

// get channel status, close was close if return a false
func (counter *Counter) GetS() bool {
	return counter.status
}

// calculate sum of counter
func (counter *Counter) Sum() int {
	counter.m.Lock()
	defer counter.m.Unlock()

	if !counter.status { // channel closed
		return counter.s
	}

	var i, n int
	for {
		if n = len(counter.ch); n == 0 {
			break
		}
		// len(counter.ch) is dynamic in for loop
		for i = 0; i < n; i++ {
			// fmt.Println(i)
			counter.s += <-counter.ch
		}
	}

	// fmt.Println(len(counter.ch), cap(counter.ch))
	for i = 0; i < n; i++ {
		counter.ch <- 0
	}

	return counter.s
}

// close channel and method Add can't useable
func (counter *Counter) Close() {
	counter.m.Lock()
	defer counter.m.Unlock()

	if !counter.status {
		return
	}

	counter.status = false
	var i, n int

	for {
		if n = len(counter.ch); n == 0 {
			break
		}
		// len(counter.ch) is dynamic in for loop
		for i = 0; i < n; i++ {
			// fmt.Println(i)
			counter.s += <-counter.ch
		}
	}

	close(counter.ch)
	counter.ch = nil
	return
}
