package rate_limit

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// rate limiter using a time array
type LimiterV3 struct {
	interval time.Duration
	vec      []time.Time
	strong   bool // count event get bucket failed
	p        int
	mu       sync.RWMutex
	exit     chan struct{}
}

func NewLimiterV3(interval time.Duration, b int, strong bool) (limiter *LimiterV3, err error) {
	if interval < time.Second || b <= 0 {
		return nil, fmt.Errorf("invalid parameter for NewLimiter")
	}

	limiter = &LimiterV3{
		interval: interval,
		vec:      make([]time.Time, b),
		strong:   strong,
		exit:     make(chan struct{}),
	}

	return
}

// (intf) New(time.Duration, int, bool) (intf, error), as golang 1.17 doesn't support generics
func (limiter *LimiterV3) New(interval time.Duration, b int, strong bool) (*LimiterV3, error) {
	return NewLimiterV3(interval, b, strong)
}

func (limiter *LimiterV3) next(now time.Time) (next int) {
	switch {
	case limiter.p == 0 && limiter.vec[0].IsZero():
		next = 0
	case limiter.p < len(limiter.vec)-1:
		next = limiter.p + 1
	default:
		next = 0
	}

	return next
}

func (limiter *LimiterV3) allow(now time.Time) (ok bool) {
	if limiter.vec[limiter.p].After(now) {
		now = time.Now()
	}

	limiter.mu.RLock()

	next := limiter.next(now)
	ok = now.Sub(limiter.vec[next]) > limiter.interval

	if limiter.strong || ok {
		limiter.mu.RUnlock()
		limiter.mu.Lock()
		limiter.p, limiter.vec[next] = next, now
		limiter.mu.Unlock()
	} else {
		limiter.mu.RUnlock()
	}

	return ok
}

func (limiter *LimiterV3) Allow(now time.Time) (ok bool) {
	select {
	case <-limiter.exit:
		return false
	default:
	}

	return limiter.allow(now)
}

func (limiter *LimiterV3) AllowWithContext(ctx context.Context, now time.Time) (ok bool) {
	select {
	case <-limiter.exit:
		return false
	case <-ctx.Done(): // allow context like a timeout
		return false
	default:
	}

	return limiter.allow(now)
}

func (limiter *LimiterV3) Last() time.Time {
	return limiter.vec[limiter.p]
}

func (limiter *LimiterV3) Stop() {
	close(limiter.exit)
}
