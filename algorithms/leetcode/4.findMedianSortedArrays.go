package leetcode

import (
	"sort"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	numx := append(nums1, nums2...)

	sort.Ints(numx)
	n := len(numx)
	if n%2 == 1 {
		return float64(numx[n/2])
	}

	return float64(numx[n/2]+numx[n/2-1]) / 2
}
