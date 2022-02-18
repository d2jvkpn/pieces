/**
 * Two Sum Input Array Is Sorted
 */
package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSumSorted([]int{1, 2, 3, 4, 5}, 7))
	fmt.Println(twoSumSorted([]int{1, 3, 7, 9, 11}, 13))
	fmt.Println(twoSumSorted([]int{1, 7, 8, 9}, 16))
}

func twoSumSorted(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}

	for i, j := 0, len(nums)-1; i < j; {
		println("~~~", i, j)
		sum := nums[i] + nums[j]

		switch {
		case sum == target:
			return []int{i, j}
		case sum < target:
			i++
		default:
			j--
		}
	}

	return nil
}
