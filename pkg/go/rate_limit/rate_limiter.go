package rate_limit

import (
	"fmt"
	"sync"
	"time"
)

type Limiter interface {
	Allow(time.Time) bool
	// AllowWithContext(context.Context, time.Time) bool
	Stop()
	Last() time.Time
}

type NewLimiter = func(time.Duration, int) Limiter

type RateLimiter struct {
	b int
	// dur      time.Duration // clear key durarion
	interval   time.Duration // rate limit durarion
	newLimiter NewLimiter
	mu         *sync.RWMutex
	ticker     *time.Ticker
	mp         map[string]Limiter
	exit       chan struct{}
}

func NewRateLimiter(secs int64, b int, newLimiter NewLimiter) (rl *RateLimiter, err error) {
	if secs < 1 || b < 1 || newLimiter == nil {
		return nil, fmt.Errorf("invlaid parameter for RateLimiter")
	}

	interval := time.Second * time.Duration(secs)

	rl = &RateLimiter{
		b:          b,
		interval:   interval,
		newLimiter: newLimiter,
		mu:         new(sync.RWMutex),
		ticker:     time.NewTicker(RATELIMITER_ClearEveryN * interval),
		mp:         make(map[string]Limiter, 100),
		exit:       make(chan struct{}),
	}

	go func() {
		for {
			select {
			case <-rl.exit:
				// fmt.Println("!!! RateLimiter.ticker runtime closed")
				return
			case <-rl.ticker.C:
				// fmt.Println("~~~ RateLimiter.ticker", rfc3339now())
			}

			now := time.Now()
			rl.mu.Lock()
			for key, limiter := range rl.mp {
				dur := now.Sub(limiter.Last())
				if dur > RATELIMITER_ClearEveryN*rl.interval {
					// fmt.Println(rfc3339now(), "RateLimiter drop key:", key, dur)
					limiter.Stop()
					delete(rl.mp, key)
				}
			}
			rl.mu.Unlock()
		}
	}()

	return
}

func (rl *RateLimiter) GetLimiter(key string) (limiter Limiter) {
	var ok bool
	rl.mu.RLock()

	if limiter, ok = rl.mp[key]; ok {
		rl.mu.RUnlock()
		return limiter
	}

	rl.mu.RUnlock()
	rl.mu.Lock()

	if limiter, ok = rl.mp[key]; !ok {
		limiter = rl.newLimiter(rl.interval, rl.b)
		rl.mp[key] = limiter
	}
	rl.mu.Unlock()

	return limiter
}

func (rl *RateLimiter) Metrics() (time.Duration, int, int) {
	return rl.interval, rl.b, len(rl.mp)
}

func (rl *RateLimiter) Allow(key string) (ok bool) {
	return rl.GetLimiter(key).Allow(time.Now())
}

func (rl *RateLimiter) Stop() {
	rl.ticker.Stop()
	close(rl.exit)

	rl.mu.Lock()
	for key, limiter := range rl.mp {
		limiter.Stop()
		delete(rl.mp, key)
	}
	rl.mu.Unlock()
}
