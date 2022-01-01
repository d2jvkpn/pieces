package main

import (
	"fmt"
)

func main() {
	vec := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := vec[:3]
	fmt.Printf("vec=%#v, slice=%#v, cap(slice)=%d\n", vec, slice, cap(slice))

	slice = slice[:6]
	fmt.Printf("vec=%#v, slice=%#v, cap(slice)=%d\n", vec, slice, cap(slice))

	arr := (*[3]int)(vec)
	fmt.Printf("arr=%#[1]v, %[1]T\n", arr)
	arr[0] = -1
	vec[2] = -2

	fmt.Printf("vec=%#v, arr=%#v\n", vec, arr)
	// vec=[]int{-1, 1, -2, 3, 4, 5, 6, 7, 8, 9}, arr=&[3]int{-1, 1, -2}

	// arr2 := (*[20]int)(vec) // panic
	// fmt.Printf("arr2=%#v\n", arr2)
}
