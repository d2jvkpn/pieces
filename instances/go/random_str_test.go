package main

import (
	"fmt"
	"testing"
)

func TestRandomStr(t *testing.T) {
	for i := 1; i < 16; i++ {
		fmt.Println(">>>", i)
		fmt.Println(RandomStr(16, i))
	}

	fmt.Println(RandomStr(16, 0b1010))
}
