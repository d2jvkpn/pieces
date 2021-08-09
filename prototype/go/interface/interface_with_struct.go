package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")

	x := NewX("Rover")
	x.Echo()

	intf := NewIntfFromX("Rover2")

	intf.Echo()

	var intf2 Intf
	var emp interface{}

	emp = x

	intf2 = emp.(Intf)
	intf2.Echo()
}

type Intf interface {
	Echo()
}

type X struct {
	Name string
}

func NewX(name string) (x *X) {
	return &X{Name: name}
}

func (x *X) Echo() {
	fmt.Println(x.Name)
}

func NewIntfFromX(name string) Intf {
	return NewX(name)
}
