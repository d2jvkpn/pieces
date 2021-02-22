package algorithms

import (
	"sort"
	"strconv"
	"strings"
)

func SliceToMaxNumber(slice []int) string {
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
