package main

import (
	"fmt"
	"time"
)

func main() {
	d := NewD()
	fmt.Println(">>> 1 main NewD():", d)
	time.Sleep(2 * time.Second)
	fmt.Println(">>> 3 main NewD():", d)
}

type D struct {
	V int64
}

func NewD() (d D) {
	d = D{V: 1}
	go func() {
		time.Sleep(1 * time.Second)
		d.V = 2
		fmt.Println(">>> 2 NewD():", d)
	}()
	return d
}

func NewD1() (d D) {
	return D{V: 1}
}

func NewD2() (d *D) {
	return &D{V: 1}
}
