package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	fmt.Printf("%v -> %s\n", strs, LongestCommonPrefix(strs))

	strs = []string{"dog", "racecar", "car"}
	fmt.Printf("%v -> %s\n", strs, LongestCommonPrefix(strs))

	strs = []string{"", "b"}
	fmt.Printf("%v -> %s\n", strs, LongestCommonPrefix(strs))
}
