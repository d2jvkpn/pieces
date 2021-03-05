package explore

import (
	"fmt"
)

func SelectSort(slice []int) (n int) {
	for i := 0; i < len(slice)-1; i++ {
		z := i
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[i] {
				z = j
			}
		}

		if z != i {
			n++
			slice[i], slice[z] = slice[z], slice[i]
		}
	}

	return
}

func InstSelectSort() {
	fmt.Println(">>> InstSelectSort:")
	slice := []int{14, 33, 10, 27, 19, 35, 42, 44, 18}
	fmt.Printf("    slice = %v\n", slice)

	n := SelectSort(slice)
	fmt.Printf("    n = %d, slice = %v\n", n, slice)
}
