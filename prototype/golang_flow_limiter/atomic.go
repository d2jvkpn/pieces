package main

import (
	"fmt"
	"sync/atomic"
)

type Counter struct {
	n int64
}

func NewCounter() (counter *Counter) {
	return &Counter{n: new(int64)}
}

func (counter *Counter) Delta(n int64) (out int64) {
	return atomic.AddInt64(counter.n, n)
}

func (counter *Counter) Value() (out int64) {
	return atomic.LoadInt64(counter.n)
}
