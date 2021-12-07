package interview

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func SliceToMaxnum(slice []int) string {
	if len(slice) == 0 {
		return "0"
	}

	strs := make([]string, len(slice))
	for i := range slice {
		strs[i] = strconv.Itoa(slice[i])
	}

	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})

	return strings.Join(strs, "")
}

func InstSliceToMaxnum() {
	fmt.Println(">>> InstSliceToMaxnum:")
	slice := []int{1, 4, 30, 34, 301, 9, 5}
	fmt.Printf("    slice = %v\n", slice)
	// expect 95434303011

	fmt.Printf("    out = %s\n", SliceToMaxnum(slice))
}
