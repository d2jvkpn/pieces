package rate_limit

import (
	"fmt"
	"testing"
	"time"
)

func TestLimiterV2_t1(t *testing.T) {
	limiter, _ := NewLimiterV2(10*time.Second, 2)

	for i := 0; i < 20; i++ {
		now := time.Now()
		fmt.Printf("~~~ %s, %t\n", now.Format(time.RFC3339), limiter.Allow(now))
		time.Sleep(time.Second)
	}
}
