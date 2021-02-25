package main

import (
	"fmt"
)

func main() {
	InstDefer1()
	InstDefer2()
}

// instance 1
type Data struct{}

func (d *Data) Echo(str string) *Data {
	fmt.Print(str)
	return d
}

func InstDefer1() {
	var d Data

	defer d.Echo("1").Echo("2\n")
	fmt.Print("3")
	// output: 132
}

// instance 2
func InstDefer2() {
	defer1 := func(i int) (t int) {
		t = i
		defer func() {
			t += 3
		}()
		return t
	}

	defer2 := func(i int) int {
		t := i
		defer func() {
			t += 3
		}()
		return t
	}

	defer3 := func(i int) (t int) {
		defer func() {
			t += i
		}()
		return 2
	}

	fmt.Print(defer1(1))
	fmt.Print(defer2(1))
	fmt.Print(defer3(1), "\n")
	// output: 413
}
