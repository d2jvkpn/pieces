package leetcode

import (
	"fmt"
)

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	min := 0
	for i := range strs {
		if i == 0 || len(strs[i]) < min {
			min = len(strs[i])
		}
	}

OUT:
	for i := 0; i < min; i++ {
		for j := range strs {
			if j < len(strs)-1 && strs[j][i] != strs[j+1][i] {
				min = i
				break OUT
			}
		}
	}

	return strs[0][:min]
}

func InstLongestCommonPrefix() {
	fmt.Println(">>> InstLongestCommonPrefix:")

	strs := []string{"flower", "flow", "flight"}
	fmt.Printf("    strs = %v,  prefix = %s\n", strs, LongestCommonPrefix(strs))

	strs = []string{"dog", "racecar", "car"}
	fmt.Printf("    strs = %v,  prefix = %s\n", strs, LongestCommonPrefix(strs))

	strs = []string{"", "b"}
	fmt.Printf("    strs = %v,  prefix = %s\n", strs, LongestCommonPrefix(strs))
}
