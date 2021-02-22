package algorithms

import (
	"fmt"
	"testing"
)

func TestSliceToMaxNumber(t *testing.T) {
	slice := []int{1, 4, 30, 301, 9, 5}

	fmt.Printf("%v -> %s\n", slice, MaxConcatNum(slice))
}
