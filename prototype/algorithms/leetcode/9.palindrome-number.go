package leetcode

import (
	"fmt"
	"strconv"
)

func PalindromeNumber(num int) bool {
	str := strconv.Itoa(num)

	for i := 0; i <= len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}

	return true
}

func InstPalindromeNumber() {
	fmt.Println(">>> InstPalindromeNumber:")
	fmt.Printf("    number = %d,  isPalindromeNumber = %t\n", 12421, PalindromeNumber(12421))
	fmt.Printf("    number = %d,  isPalindromeNumber = %t\n", 124, PalindromeNumber(124))
	fmt.Printf("    number = %d,  isPalindromeNumber = %t\n", -12421, PalindromeNumber(-12421))
}
