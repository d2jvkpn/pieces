package interview

import (
	"fmt"
	"testing"
)

func TestSortPN(t *testing.T) {
	arr := []int{4, 1, 3, 2, -3, -1, -4, -2}
	// expect 1234, -1,-2,-3,-4

	SortPN(arr)
	fmt.Println(arr)
}
