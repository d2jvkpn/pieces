package explore

import (
	"fmt"
)

func QuickSort(slice []int) (out []int) {
	if len(slice) < 2 {
		return slice
	}

	s1, s2 := make([]int, 0, len(slice)/2), make([]int, 0, len(slice)/2)

	for i := 1; i < len(slice); i++ {
		if slice[i] < slice[0] {
			s1 = append(s1, slice[i])
		} else {
			s2 = append(s2, slice[i])
		}
	}
	fmt.Printf("    s1 = %v, s2 = %v\n", s1, s2)

	out = make([]int, 0, len(slice))
	out = append(QuickSort(s1), slice[0]) // recursion
	out = append(out, QuickSort(s2)...)   // recursion

	return
}

func InstQuickSort() {
	fmt.Println(">>> InstQuickSort:")
	slice := []int{14, 33, 10, 27, 19, 35, 42, 44, 18}
	fmt.Printf("    slice = %v\n", slice)

	out := QuickSort(slice)
	fmt.Printf("    out = %v\n", out)
}
