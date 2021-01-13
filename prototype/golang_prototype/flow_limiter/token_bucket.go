package main

import (
	"fmt"
	"sync/atomic"
)

type Data interface{}

type TokenBucket struct {
	count *uint64
	limit uint64
}

func NewTokenBucket(n uint64) (lim *TokenBucket) {
	return &TokenBucket{count: new(uint64), limit: n}
}

func (lim *TokenBucket) String() string {
	return fmt.Sprintf("count: %d, limit: %d", atomic.LoadUint64(lim.count), lim.limit)
}

func (lim *TokenBucket) Handle(d Data, overflow, do func(Data)) {
	if atomic.LoadUint64(lim.count) >= lim.limit {
		overflow(d)
	} else {
		atomic.AddUint64(lim.count, 1)
		do(d)
		var i int64 = -1
		atomic.AddUint64(lim.count, uint64(i))
	}
}
