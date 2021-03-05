package explore

import (
	"fmt"
)

func InsertSort(slice []int) (n int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if slice[j-1] > slice[j] {
				slice[j-1], slice[j] = slice[j], slice[j-1]
				n++
			}
		}
	}

	return n
}

func InstInsertSort() {
	fmt.Println(">>> InsertSort:")
	slice := []int{56, 1, 7, 2, 4, 9, 19, 16, 32, 30}
	fmt.Printf("    slice = %v\n", slice)

	n := InsertSort(slice)
	fmt.Printf("    n = %d, slice = %v\n", n, slice)
}
