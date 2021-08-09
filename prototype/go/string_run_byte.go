package main

import (
	"fmt"
)

func main() {
	s := "foö" // Unicode: f=0x66 o=0x6F ö=0xC3B6
	r := []rune(s)
	b := []byte(s)

	fmt.Printf("%T %v\n", s, s) // "string foö"
	fmt.Printf("%T %v\n", r, r) // "[]int32 [102 111 246]"
	fmt.Printf("%T %v\n", b, b) // "[]uint8 [102 111 195 182]"
}
