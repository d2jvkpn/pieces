package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

func main() {
	var secs int64

	flag.Int64Var(&secs, "secs", 4, "time.After seconds")
	flag.Parse()

	ch := make(chan int)
	go call(ch)

	select {
	case v := <-ch:
		fmt.Printf("    got %d from ch\n", v)
	case <-time.After(time.Duration(secs) * time.Second):
		fmt.Printf("    time out\n")
	}

	time.Sleep(10 * time.Second)
	fmt.Printf(">>> NumGoroutine: %d\n", runtime.NumGoroutine())
}

func call(ch chan<- int) {
	time.Sleep(5 * time.Second)
	fmt.Println("    call done")

	ch <- 1
	fmt.Println("    call return")
}
