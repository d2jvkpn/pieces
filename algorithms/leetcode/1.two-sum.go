package leetcode

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
