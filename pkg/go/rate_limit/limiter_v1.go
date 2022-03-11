package rate_limit

import (
	"fmt"
	"time"
)

// rate limiter using a time array
type LimiterV1 struct {
	interval time.Duration
	vec      []time.Time
	strong   bool
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

func (limiter *LimiterV1) next(now time.Time) (old time.Time) {
	// fmt.Println("-->", limiter.p, limiter.vec)
	switch {
	case limiter.p == 0 && limiter.vec[limiter.p].IsZero():
		// the first value
		// fmt.Println("...0")
	case limiter.p < len(limiter.vec)-1:
		// fmt.Println("...1")
		limiter.p++
	default:
		// fmt.Println("...-1")
		limiter.p = 0
	}
	old = limiter.vec[limiter.p] // extract old value

	return old
}

func (limiter *LimiterV1) Allow() (ok bool) {
	select {
	case <-limiter.ch:
		return false
	case limiter.ch <- struct{}{}:
	}
	defer func() { <-limiter.ch }()

	now := time.Now()
	ok = now.Sub(limiter.next(now)) > limiter.interval

	if limiter.strong || ok {
		limiter.vec[limiter.p] = now
	}

	return
}

func (limiter *LimiterV1) Stop() {
	close(limiter.exit)
}
