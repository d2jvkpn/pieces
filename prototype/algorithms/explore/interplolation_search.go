package explore

import (
	"fmt"
)

func IntеrpоlаtiоnSеаrch(slice []int, target int) int {
	n, m, idx := 0, len(slice)-1, 0
	if len(slice) == 0 || target < slice[n] || target > slice[m] {
		return -1
	}

	for {
		vn, vm := slice[n], slice[m]
		idx2 := n + 1 + (m-n)*(target-vn)/(vm-vn)
		if idx == idx2 {
			break
		}
		idx = idx2
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

func InstIntеrpоlаtiоnSеаrch() {
	fmt.Println(">>> InstIntеrpоlаtiоnSеаrch:")
	slice := []int{1, 4, 7, 9, 10, 14, 17, 20, 27, 31}
	fmt.Printf("    slice = %v\n", slice)

	var target int

	target = 17
	fmt.Printf("    target = %d, result = %d\n", target, IntеrpоlаtiоnSеаrch(slice, target))

	target = 7
	fmt.Printf("    target = %d, result = %d\n", target, IntеrpоlаtiоnSеаrch(slice, target))

	target = 100
	fmt.Printf("    target = %d, result = %d\n", target, IntеrpоlаtiоnSеаrch(slice, target))
}
