package main

import (
	"fmt"
)

func main() {
	mp := map[string]string{"a": "A"}
	fmt.Println(">>> 1", mp)

	call1(mp)
	fmt.Println(">>> 3", mp)

	slice := []int64{1, 2, 3}
	fmt.Println(">>> 4", slice)

	call2(slice)
	fmt.Println(">>> 6", slice)
}

func call1(mp map[string]string) {
	mp = nil
	fmt.Println(">>> 2", mp)
}

func call2(slice []int64) {
	slice = nil
	fmt.Println(">>> 5", slice)
}
