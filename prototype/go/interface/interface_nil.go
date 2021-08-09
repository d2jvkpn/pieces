package main

import (
	"fmt"
)

func main() {
	var i int = 64

	x := (*Intf)(nil)
	fmt.Printf("x: %[1]T, %[1]v, is nil: %[2]t\n", x, x == nil)

	s := (string)(i) // same as string(i)
	fmt.Printf("%[1]T, %[1]v\n", s)

	d := (*D)(nil)
	fmt.Printf("d: %#v, type: %T, is nil: %t\n", d, d, d == nil)
}

type Intf interface {
	Echo()
}

type D struct {
	A int64
}
