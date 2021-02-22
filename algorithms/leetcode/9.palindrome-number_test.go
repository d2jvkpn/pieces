package leetcode

import (
	"fmt"
	"testing"
)

func TestPalindromeNumber(t *testing.T) {
	fmt.Printf("%d -> %t\n", 12421, PalindromeNumber(12421))
	fmt.Printf("%d -> %t\n", 124, PalindromeNumber(124))
	fmt.Printf("%d -> %t\n", -12421, PalindromeNumber(-12421))
}
