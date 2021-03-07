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

	var target int
	target = 17
	fmt.Printf("    target = %d, result = %d\n", target, BinarySearch(slice, target))

	target = 7
	fmt.Printf("    target = %d, result = %d\n", target, BinarySearch(slice, target))

	target = 100
	fmt.Printf("    target = %d, result = %d\n", target, BinarySearch(slice, target))
}

func InstBinarySearch2() {
	fmt.Println(">>> InstBinarySearch2:")
	slice := []int{1, 4}
	fmt.Printf("    slice = %v\n", slice)

	var target int

	target = 1
	fmt.Printf("    target = %d, result = %d\n", target, BinarySearch(slice, target))

	target = 4
	fmt.Printf("    target = %d, result = %d\n", target, BinarySearch(slice, target))

	target = 2
	fmt.Printf("    target = %d, result = %d\n", target, BinarySearch(slice, target))
}

func InstBinarySearch3() {
	fmt.Println(">>> InstBinarySearch3:")
	slice := []int{1, 2}
	fmt.Printf("    slice = %v\n", slice)

	var target int

	target = 1
	fmt.Printf("    target = %d, result = %d\n", target, BinarySearch(slice, target))

	target = 2
	fmt.Printf("    target = %d, result = %d\n", target, BinarySearch(slice, target))

	target = -1
	fmt.Printf("    target = %d, result = %d\n", target, BinarySearch(slice, target))
}
