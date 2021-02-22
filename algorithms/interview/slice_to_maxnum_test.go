package interview

import (
	"fmt"
	"testing"
)

func TestSliceToMaxNumber(t *testing.T) {
	slice := []int{1, 4, 30, 34, 301, 9, 5}
	// expect

	fmt.Printf("%v -> %s\n", slice, MaxConcatNum(slice))
}
