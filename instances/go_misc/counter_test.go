package main

import (
	"fmt"
	"math/rand"
	// "runtime"
	"log"
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	var s int
	// runtime.GOMAXPROCS(2)
	c, _ := NewCounter(12)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	t0 := time.Now()
	// concurrency number
	for i := 0; i < 1e8; i++ {
		if i%1e4 == 1000 {
			log.Println(i)
		}
		k := r.Intn(100)
		s += k

		go c.Add(k) // go c.Add(k) will cause high cpu and memory usage
	}

	fmt.Println(time.Now().Sub(t0))
	n := c.Sum()

	fmt.Printf("counter=%d, total=%d\n", n, s)
	if n != s {
		t.Errorf("counter and total are not equal\n")
	}
}

func TestCounterCmp(t *testing.T) {
	var s int
	var m sync.Mutex

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	t0 := time.Now()
	for i := 0; i < 1e8; i++ {
		k := r.Intn(100)
		m.Lock()
		s += k
		m.Unlock()
	}

	fmt.Println(time.Now().Sub(t0))
}
