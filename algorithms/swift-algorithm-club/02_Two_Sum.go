/**
 * Two Sum
 */
package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5}, 6))
}

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}

	index := make(map[int]int, len(nums))

	for i, v := range nums {
		if p, ok := index[target-v]; ok {
			return []int{v, p}
		}
		index[v] = i
	}

	return nil
}
