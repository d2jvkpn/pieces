package main

import (
	"fmt"
)

func main() {
	v := 1
	fmt.Println(">>> ", v)

	call1(&v)
	fmt.Println(">>> ", v)

	call2(&v)
	fmt.Println(">>> ", v)
}

func call1(p *int) {
	*p = 2
}

func call2(p *int) {
	x := 32
	p = &x
}
