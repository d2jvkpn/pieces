package main

import (
	"fmt"
)

func main() {
	call02()
}

func call01() (n int) {
	fmt.Println("call01")
	return 42
}

func call02() (n int) {
	fmt.Println("call02")
	defer func() {
		fmt.Println("defer in call02")
	}()

	return call01()
}
