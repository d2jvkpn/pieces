package main

import (
	"fmt"
	"time"
)

func main() {
	d := NewD()
	fmt.Println(d)
	time.Sleep(3 * time.Second)
	fmt.Println(d)
}

type D struct {
	V int64
}

func NewD() (d D) {
	d = D{V: 1}
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("ok")
		d.V = 2
	}()
	return d
}
