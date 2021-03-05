package leetcode

import (
	"fmt"
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

func InstFindMedianSortedArrays() {
	fmt.Println(">>> InstFindMedianSortedArrays:")
	slice1 := []int{2, 9, 19, 30, 12, 6}
	slice2 := []int{2, 100, 9, 3, 10, 82}
	fmt.Printf("    slice1 = %v, slice2 = %v\n", slice1, slice2)

	out := FindMedianSortedArrays(slice1, slice2)
	fmt.Printf("    out = %v\n", out)
}
