package leetcode

import (
	"fmt"
)

func TwoSum(nums []int, target int) []int {
	mp := make(map[int]int, len(nums))
	for i := range nums {
		mp[nums[i]] = i
	}

	for i := range nums {
		j, ok := mp[target-nums[i]]
		if i != j && ok {
			return []int{i, j}
		}
	}

	return nil
}

func InstTwoSum() {
	fmt.Println(">>> InstTwoSum")
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 10
	fmt.Printf("    slice = %v, target = %d\n", slice, target)

	out := TwoSum(slice, target)
	fmt.Printf("    out = %v\n", out)
}
