package interview

import (
	"fmt"
	"sort"
)

func SortPN(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		switch {
		case arr[i]*arr[j] < 0:
			return arr[i] > arr[j]
		case arr[i] < 0:
			return arr[i] > arr[j]
		default:
			return arr[i] < arr[j]
		}
	})
}

func InstSortPN() {
	arr := []int{4, 1, 3, 2, -3, -1, -4, -2}
	// expect 1, 2, 3, 4, -1,-2,-3,-4

	SortPN(arr)
	fmt.Println(arr)
}
