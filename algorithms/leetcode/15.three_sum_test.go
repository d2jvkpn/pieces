package leetcode

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	ints := []int{-1, 0, 1, 2, -1, -4}
	fmt.Printf("%v -> %v\n", ints, ThreeSum(ints))
}

func TestThreeSum2(t *testing.T) {
	ints := []int{-1, 0, 1, 2, -1, -4}
	fmt.Printf("%v -> %v\n", ints, ThreeSum2(ints))
}

func TestThreeSum3(t *testing.T) {
	ints := []int{-1, 0, 1, 2, -1, -4}
	fmt.Printf("%v -> %v\n", ints, ThreeSum3(ints))
}
