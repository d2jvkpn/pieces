package explore

import (
	"fmt"
)

func BinarySearch(slice []int, target int) int {
	fmt.Printf(">>> BinarySearch:\n    slice -> %v, target -> %d\n", slice, target)

	n, m := 0, len(slice)-1
	if len(slice) == 0 || target < slice[n] || target > slice[m] {
		return -1
	}

	for {
		/*
			v := (float64(n) + float64(m))/2
			idx := int(v) // idx := int(math.Ceil(v))
		*/
		idx := (n + m) / 2
		fmt.Printf("    n = %d, m = %d, idx = %d\n", n, m, idx)

		switch {
		case slice[idx] == target:
			return idx
		// n = 0 && m = 1 => idx = 0, check target == slice[1]
		case slice[idx+1] == target:
			return idx + 1
		case slice[idx] > target:
			m = idx
		case slice[idx] < target:
			n = idx
		}

		if idx == (n+m)/2 {
			break
		}
	}

	return -1
}

func InstBinarySearch1() {
	fmt.Println(">>> InstBinarySearch1:")
	slice := []int{1, 4, 7, 9, 10, 14, 17, 20, 27, 31}
	fmt.Printf("    slice = %v\n", slice)

	fmt.Println("    result =", BinarySearch(slice, 17))
	fmt.Println("    result =", BinarySearch(slice, 7))
	fmt.Println("    result =", BinarySearch(slice, 100))
}

func InstBinarySearch2() {
	fmt.Println(">>> InstBinarySearch2:")
	slice := []int{1, 4}
	fmt.Printf("    slice = %v\n", slice)

	fmt.Println("    result =", BinarySearch(slice, 1))
	fmt.Println("    result =", BinarySearch(slice, 4))
	fmt.Println("    result =", BinarySearch(slice, 2))
}

func InstBinarySearch3() {
	fmt.Println(">>> InstBinarySearch3:")
	slice := []int{1, 2}
	fmt.Printf("    slice = %v\n", slice)

	fmt.Println("    result =", BinarySearch(slice, 1))
	fmt.Println("    result =", BinarySearch(slice, 2))
	fmt.Println("    result =", BinarySearch(slice, -1))
}
