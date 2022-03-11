package rate_limit

import (
	"fmt"
	"time"
)

// rate limiter using a time array
type LimiterV1 struct {
	interval time.Duration
	vec      []time.Time
	strong   bool // count event get bucket failed
	p        int
	ch       chan struct{}
	exit     chan bool
}

func NewLimiterV1(interval time.Duration, b int, strong bool) (limiter *LimiterV1, err error) {
	if interval <= 0 || b <= 0 {
		return nil, fmt.Errorf("invalid parameter for NewLimiter")
	}

	limiter = &LimiterV1{
		interval: interval,
		vec:      make([]time.Time, b),
		strong:   strong,
		ch:       make(chan struct{}, 1),
		exit:     make(chan bool),
	}

	return
}

func (limiter *LimiterV1) pNext(now time.Time) (oldest time.Time) {
	switch {
	case limiter.p == 0 && limiter.vec[limiter.p].IsZero():
	case limiter.p < len(limiter.vec)-1:
		limiter.p++
	default:
		limiter.p = 0
	}

	oldest = limiter.vec[limiter.p] // extract oldest value

	return oldest
}

func (limiter *LimiterV1) pBack() {
	switch {
	case limiter.p == 0 && limiter.vec[limiter.p].IsZero():
	case limiter.p > 0:
		limiter.p--
	default:
		limiter.p = len(limiter.vec) - 1
	}
}

func (limiter *LimiterV1) Allow() (ok bool) {
	select {
	case <-limiter.ch:
		return false
	case limiter.ch <- struct{}{}:
	}
	defer func() { <-limiter.ch }()

	now := time.Now()
	ok = now.Sub(limiter.pNext(now)) > limiter.interval

	if limiter.strong || ok {
		limiter.vec[limiter.p] = now
	} else {
		limiter.pBack()
	}

	return
}

func (limiter *LimiterV1) Stop() {
	close(limiter.exit)
}
