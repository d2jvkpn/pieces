package explore

import (
	"fmt"
)

func BinarySearch(slice []int, target int) (idx int) {
	fmt.Printf(">>> BinarySearch:\n    slice -> %v, target -> %d\n", slice, target)

	n, m := 0, len(slice)-1
	if len(slice) == 0 || target < slice[n] || target > slice[m] {
		return -1
	}

	for {
		if idx == (n+m)/2 {
			break
		}
		idx = (n + m) / 2
		fmt.Printf("    n = %d, m = %d, idx = %d\n", n, m, idx)

		switch {
		case slice[idx] == target:
			return idx
		case slice[idx] > target:
			m = idx
		case slice[idx] < target:
			n = idx
		}
	}

	return -1
}

func InstBinarySearch() {
	slice := []int{
		1, 4, 7, 9, 10,
		14, 17, 20, 27, 31,
	}

	fmt.Println("    result =", BinarySearch(slice, 17))
	fmt.Println("    result =", BinarySearch(slice, 7))
	fmt.Println("    result =", BinarySearch(slice, 100))
}
