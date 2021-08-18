package main

import (
	"fmt"
)

func main() {
	///
	s := make([]byte, 2, 4)
	s0 := (*[0]byte)(s)
	fmt.Println("s0 == nil:", s0 == nil)

	s1 := (*[1]byte)(s[1:])
	fmt.Println("&s1[0] == &s[1]:", &s1[0] == &s[1])

	s2 := (*[2]byte)(s)
	fmt.Println("&s2[1] == &s1[0]:", &s2[1] == &s1[0])

	// s4 := (*[4]byte)(s)     //!! panics: len([4]byte) > len(s)

	///
	var t []string
	t0 := (*[0]string)(t) // t0 == nil
	fmt.Println("t0 == nil:", t0 == nil)

	// t1 := (*[1]string)(t) //!! panics: len([1]string) > len(t)

	///
	u := make([]byte, 0)
	u0 := (*[0]byte)(u)
	fmt.Println("u0 == nil: ", u0 == nil)
}
