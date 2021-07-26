package main

import (
	"fmt"
	"time"
)

func main() {
	d := NewD()
	fmt.Printf(">>> 1 main NewD(): %[1]v, %[1]p\n", &d)
	time.Sleep(2 * time.Second)
	fmt.Printf(">>> 3 main NewD(): %[1]v, %[1]p\n", &d)

	d2 := NewD2()
	fmt.Printf(">>> 3 main NewD2(): %[1]v, %[1]p\n", d2)
}

type D struct {
	V int64
}

func NewD() (d D) {
	d = D{V: 1}
	go func() {
		time.Sleep(1 * time.Second)
		d.V = 2
		fmt.Printf(">>> 2 NewD(): %[1]v, %[1]p\n", &d)
	}()
	return d
}

func NewD1() (d D) {
	return D{V: 1}
}

func NewD2() (d *D) {
	d = &D{V: 1}
	fmt.Printf(">>> NewD2(): %[1]v, %[1]p\n", d)
	return d
}
