package leetcode

import (
	"fmt"
	"sort"
)

func ThreeSum1(nums []int) (out [][]int) {
	out = make([][]int, 0)
	mp := make(map[[3]int]bool)

	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			for z := j + 1; z < len(nums); z++ {
				if nums[i]+nums[j]+nums[z] != 0 {
					continue
				}
				slice := []int{nums[i], nums[j], nums[z]}
				sort.Ints(slice)
				arr := [3]int{slice[0], slice[1], slice[2]}
				if mp[arr] {
					continue
				}

				out = append(out, slice)
				mp[arr] = true
			}
		}
	}

	return
}

func ThreeSum2(nums []int) (out [][]int) {
	out = make([][]int, 0)
	mp := make(map[int]int)
	mp2 := make(map[[3]int]bool, len(out))
	arr := [3]int{0, 0, 0}

	for i := range nums {
		mp[nums[i]]++
	}

	add2out := func(slice []int) {
		sort.Ints(slice)
		arr[0], arr[1], arr[2] = slice[0], slice[1], slice[2]
		if mp2[arr] {
			return
		}

		out = append(out, slice)
		mp2[arr] = true
	}

	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			mp[nums[i]]--
			mp[nums[j]]--
			v := 0 - nums[i] - nums[j]
			if mp[v] > 0 {
				add2out([]int{nums[i], nums[j], v})
			}
			mp[nums[i]]++
			mp[nums[j]]++
		}
	}

	return
}

func ThreeSum3(nums []int) (out [][]int) {
	out = make([][]int, 0)
	mp := make(map[int]int)
	mp2 := make(map[[3]int]bool)
	arr := [3]int{0, 0, 0}

	for i := range nums {
		mp[nums[i]]++
	}

	add2out := func(slice []int) {
		sort.Ints(slice)
		arr[0], arr[1], arr[2] = slice[0], slice[1], slice[2]
		if mp2[arr] {
			return
		}

		out = append(out, slice)
		mp2[arr] = true
	}

	for k := range mp {
		mp[k]--
		for v := range mp {
			if mp[v] == 0 {
				continue
			}
			mp[v]--
			if mp[0-k-v] > 0 {
				add2out([]int{k, v, 0 - k - v})
			}
			mp[v]++
		}
		mp[k]++
	}

	return
}

func InstThreeSum() {
	fmt.Println(">>> InstThreeSum:")
	slice := []int{-1, 0, 1, 2, -1, -4}
	fmt.Printf("    slice = %v\n", slice)

	fmt.Printf("    ThreeSum1 out = %v\n", ThreeSum1(slice))
	fmt.Printf("    ThreeSum2 out = %v\n", ThreeSum2(slice))
	fmt.Printf("    ThreeSum3 out = %v\n", ThreeSum3(slice))
}
