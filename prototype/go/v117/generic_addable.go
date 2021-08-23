package main

import (
	"fmt"
)

type Addable interface {
	type int, int64, float64, string
}

func add[T Addable](a, b T) T {
	return a+b
}

func main() {
	fmt.Println("add(1, 2) =", add(1, 2))
	fmt.Println(`add("1", "2") =`, add("1", "2"))
}


