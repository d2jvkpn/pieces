package letcode

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	result := FindMedianSortedArrays(
		[]int{2, 9, 19, 30, 12, 6},
		[]int{2, 100, 9, 3, 10, 82},
	)

	fmt.Println(result)
}
