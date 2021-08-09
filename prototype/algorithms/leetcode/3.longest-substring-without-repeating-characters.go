package leetcode

import (
	"fmt"
)

func LengthOfLongestSubstring(s string) int {
	mp := make(map[byte]int)
	m := 0

	for i := range s {
		if n := mp[s[i]]; n > 0 {
			for k := range mp {
				if mp[k] <= n {
					delete(mp, k)
				}
			}
		}

		mp[s[i]] = i + 1
		if len(mp) > m {
			m = len(mp)
		}
	}

	return m
}

func InstLengthOfLongestSubstring() {
	fmt.Println(">>> InstLengthOfLongestSubstring:")
	strs := "abccdefkk"
	fmt.Printf("    strs = %v\n", strs)

	out := LengthOfLongestSubstring("abccdefkk")
	fmt.Printf("    out = %v\n", out)
}
