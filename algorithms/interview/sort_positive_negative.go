package interview

import (
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
