package leetcode

import (
	"fmt"
	"testing"
)

func TestRomanToInteger(t *testing.T) {
	for _, str := range []string{"III", "IX", "LVIII"} {
		fmt.Printf("%s -> %d\n", str, RomanToInteger(str))
	}
}
