package rate_limit

import (
	"testing"
	"time"
)

// go test  -run none  -bench ^BenchmarkLimiterV3_b1$
// # 88.72 ns/op
func BenchmarkLimiterV3_b1(b *testing.B) {
	limiter, _ := NewLimiterV3(1*time.Second, 1, false)

	now := time.Now()
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			limiter.Allow(now)
		}
	})
}

// go test  -run none  -bench ^BenchmarkLimiterV3_b2$
func BenchmarkLimiterV3_b2(b *testing.B) {
	limiter, _ := NewLimiterV3(60*time.Second, 3, false)

	now := time.Now()
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			limiter.Allow(now)
		}
	})
}
