package main

import (
	"fmt"
)

func main() {
	var slice []int
	fmt.Println(">>> 1:", slice)

	call(slice)

	slice = append(slice, 10)
	fmt.Println(">>> 2:", slice)
}

func call(slice []int) {
	slice = []int{1, 2, 3}
}
