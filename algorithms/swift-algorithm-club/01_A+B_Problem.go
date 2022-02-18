/**
 *A + B Problem
 */
package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d + %d = %d\n", 100, 2, getSum(100, 2))

	fmt.Printf("(%d + %d)/2 = %d\n", 3, 5, mean(3, 5))

	fmt.Println(1|2, 1&2, 2<<1, 1^2)
}

func getSum(a, b int) int {
	for a != 0 {
		a, b = (a&b)<<1, a^b // AND, Left Shift, XOR
	}

	return b
}

////
func mean(a, b int) int {
	return a>>1 + b>>1 + a&b&1
}
