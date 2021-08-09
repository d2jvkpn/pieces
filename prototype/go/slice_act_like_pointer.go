package main

import (
	"fmt"
)

func main() {
	/// instance 1
	var slice []int
	fmt.Println(">>> 1:", slice)

	call(slice)

	slice = append(slice, 10)
	fmt.Println(">>> 2:", slice)

	/// instance 2
	x1 := make([]int, 0, 2)
	x2 := append(x1, 1, 2, 3, 4)
	fmt.Println(">>> 3:", x1, x2) // [] [1 2 3 4]

	/// instance 3
	arr := [10]int{}
	y1 := arr[5:6]
	fmt.Println(y1, len(y1), cap(y1)) // [1] 1 5

	y2 := append(y1, 1, 2, 3, 4)
	fmt.Println(">>> 4:", arr, y1, y2)
	// [0 0 0 0 0 0 1 2 3 4]
}

func call(slice []int) {
	slice = []int{1, 2, 3}
}
