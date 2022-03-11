package rate_limit

import (
	"fmt"
	"testing"
	"time"
)

func TestRateLimiterV2(t *testing.T) {
	rl, _ := NewRateLimiterV2(5, 2)
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
