package rate_limit

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	ctx := context.TODO()
	fmt.Println(ctx.Done() == nil) // true

	<-ctx.Done() // aways block
	fmt.Println("done!")
}

func TestChannel(t *testing.T) {
	c1 := make(chan struct{})
	fmt.Println(c1 == nil)
	close(c1)
	fmt.Println(c1 == nil)

	c2 := make(chan bool, 1)
	fmt.Println(c2 == nil)
	close(c2)
	fmt.Println(c2 == nil)
}

func TestLimiterV1_t1(t *testing.T) {
	limiter, _ := NewLimiterV1(10*time.Second, 3, false)

	for i := 0; i < 20; i++ {
		now := time.Now()
		fmt.Printf("~~~ %s, %t\n", rfc3339now(), limiter.Allow(now))
		time.Sleep(time.Second)
	}
}

func TestLimiterV1_t2(t *testing.T) {
	limiter, _ := NewLimiterV1(10*time.Second, 3, true)

	for i := 0; i < 20; i++ {
		now := time.Now()
		fmt.Printf("~~~ %s, %t\n", rfc3339now(), limiter.Allow(now))
		time.Sleep(time.Second)
	}
}

// go test  -run none  -bench ^BenchmarkLimiterV1_b1$
// # 938.6 ns/op
func BenchmarkLimiterV1_b1(b *testing.B) {
	limiter, _ := NewLimiterV1(1*time.Second, 1, false)

	now := time.Now()
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			limiter.Allow(now)
		}
	})
}

// go test  -run none  -bench ^BenchmarkLimiterV1_b2$
func BenchmarkLimiterV1_b2(b *testing.B) {
	limiter, _ := NewLimiterV1(60*time.Second, 3, false)

	now := time.Now()
	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			limiter.Allow(now)
		}
	})
}
