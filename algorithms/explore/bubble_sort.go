package explore

import (
	"fmt"
)

func BubbleSort(slice []int) (k, n int) {
	for {
		next := false
		for i := 0; i < len(slice)-1; i++ {
			if slice[i] > slice[i+1] {
				next = true
				n++
				slice[i], slice[i+1] = slice[i+1], slice[i]
			}
		}
		if !next {
			break
		}
		k++
	}

	return
}

func InstBubbleSort() {
	fmt.Println(">>> BubbleSort:")
	slice := []int{
		56, 1, 7, 2, 4, 9, 19, 16, 32, 30,
	}
	fmt.Printf("    slice = %v\n", slice)
	k, n := BubbleSort(slice)
	fmt.Printf("    k = %d, n = %d, slice = %v\n", k, n, slice)
}
