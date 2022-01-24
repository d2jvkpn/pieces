package main

import (
	"fmt"
)

func main() {
	var a = new(int)
	*a = 9
	call(a)
	fmt.Println(">>>", *a) // 9

	mut(a)
	fmt.Println(">>>", *a) // 11
}

func call(x *int) {
	t := 10
	x = &t
}

func mut(x *int) {
	*x = 11
}
