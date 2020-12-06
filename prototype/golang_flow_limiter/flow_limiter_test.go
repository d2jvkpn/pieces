package main

import (
	"fmt"
	"testing"
	"time"
)

func TestMain_t1(t *testing.T) {
	fl, _ := NewFlowLimiter(time.Second, 1e6)

	fmt.Println(fl.Len())
	fmt.Println(fl.Get())
	fmt.Println(fl.Get())

	fmt.Println(fl.Len())
	time.Sleep(500 * time.Millisecond)
	fmt.Println(fl.Len())

	fl.Close()
}

func TestMain_t2(t *testing.T) {
	fl, _ := NewFlowLimiter(time.Second, 1e9)

	fmt.Println(fl.Len())
	fmt.Println(fl.Get())
	fmt.Println(fl.Get())

	fmt.Println(fl.Len())
	time.Sleep(500 * time.Millisecond)
	fmt.Println(fl.Len())

	fl.Close()
}
