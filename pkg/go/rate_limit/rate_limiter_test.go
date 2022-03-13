package rate_limit

import (
	"fmt"
	"testing"
	"time"
)

// go test -run none -bench ^BenchmarkRateLimiterV1$ -count 10
// # 6417 ns/op
func BenchmarkRateLimiterV1(b *testing.B) {
	newLimiter := func(dur time.Duration, b int) Limiter {
		limiter, _ := NewLimiterV1(dur, b, true)
		return limiter
	}

	rl, _ := NewRateLimiter(5, 3, newLimiter)
	addr := "127.0.0.1"

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rl.Allow(addr)
		}
	})
}

func TestRateLimiterV2(t *testing.T) {
	newLimiter := func(dur time.Duration, b int) Limiter {
		limiter, _ := NewLimiterV2(dur, b)
		return limiter
	}

	rl, _ := NewRateLimiter(5, 3, newLimiter)
	addr := "127.0.0.1"

	task := func() {
		fmt.Printf("~~~ %s get token: %t\n", rfc3339now(), rl.Allow(addr))
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 5; j++ {
			go task()
		}
		time.Sleep(2 * time.Second)
	}

	fmt.Println("sleep", 30*time.Second)
	time.Sleep(30 * time.Second)

	for i := 0; i < 10; i++ {
		for j := 0; j < 5; j++ {
			go task()
		}
		time.Sleep(2 * time.Second)
	}
}

// go test -run none -bench ^BenchmarkRateLimiterV2$ -count 10
// # 349.0 ns/op
func BenchmarkRateLimiterV2(b *testing.B) {
	newLimiter := func(dur time.Duration, b int) Limiter {
		limiter, _ := NewLimiterV2(dur, b)
		return limiter
	}

	rl, _ := NewRateLimiter(5, 3, newLimiter)
	addr := "127.0.0.1"

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rl.Allow(addr)
		}
	})
}

// go test -run none -bench ^BenchmarkRateLimiterV3$ -count 10
// # 1725 ns/op
func BenchmarkRateLimiterV3(b *testing.B) {
	newLimiter := func(dur time.Duration, b int) Limiter {
		limiter, _ := NewLimiterV3(dur, b, true)
		return limiter
	}

	rl, _ := NewRateLimiter(5, 3, newLimiter)
	addr := "127.0.0.1"

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rl.Allow(addr)
		}
	})
}
