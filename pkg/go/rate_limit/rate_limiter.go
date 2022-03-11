package rate_limit

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiterV2 struct {
	b int
	// dur      time.Duration // clear key durarion
	interval time.Duration // rate limit durarion
	mu       *sync.RWMutex
	ticker   *time.Ticker
	mp       map[string]*LimiterV2
	exit     chan struct{}
}

func NewRateLimiterV2(secs int64, b int) (rl *RateLimiterV2, err error) {
	if secs < 1 || b < 1 {
		return nil, fmt.Errorf("invlaid parameter for RateLimiter")
	}

	interval := time.Second * time.Duration(secs)

	rl = &RateLimiterV2{
		b:        b,
		interval: interval,
		mu:       new(sync.RWMutex),
		ticker:   time.NewTicker(RATELIMITER_ClearEveryN * interval),
		mp:       make(map[string]*LimiterV2, 100),
		exit:     make(chan struct{}),
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
				dur := now.Sub(limiter.last)
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

func (rl *RateLimiterV2) GetLimiter(key string) (limiter *LimiterV2) {
	var ok bool
	rl.mu.RLock()

	if limiter, ok = rl.mp[key]; ok {
		rl.mu.RUnlock()
		return limiter
	}

	rl.mu.RUnlock()
	rl.mu.Lock()

	if limiter, ok = rl.mp[key]; !ok {
		limiter, _ = NewLimiterV2(rl.interval, rl.b)
		rl.mp[key] = limiter
	}
	rl.mu.Unlock()

	return limiter
}

func (rl *RateLimiterV2) Metrics() (time.Duration, int, int) {
	return rl.interval, rl.b, len(rl.mp)
}

func (rl *RateLimiterV2) Allow(key string) (ok bool) {
	return rl.GetLimiter(key).Allow(time.Now())
}

func (rl *RateLimiterV2) Stop() {
	rl.ticker.Stop()
	close(rl.exit)

	rl.mu.Lock()
	for key, limiter := range rl.mp {
		limiter.Stop()
		delete(rl.mp, key)
	}
	rl.mu.Unlock()
}
