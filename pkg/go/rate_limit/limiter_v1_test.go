package rate_limit

import (
	"fmt"
	"testing"
	"time"
)

func TestLimiterV1_t1(t *testing.T) {
	limiter, _ := NewLimiterV1(10*time.Second, 2, false)

	for i := 0; i < 20; i++ {
		now := time.Now()
		fmt.Printf("~~~ %s, %t\n", now.Format(time.RFC3339), limiter.Allow())
		time.Sleep(time.Second)
	}
}

func TestLimiterV1_t2(t *testing.T) {
	limiter, _ := NewLimiterV1(10*time.Second, 2, true)

	for i := 0; i < 20; i++ {
		now := time.Now()
		fmt.Printf("~~~ %s, %t\n", now.Format(time.RFC3339), limiter.Allow())
		time.Sleep(time.Second)
	}
}
