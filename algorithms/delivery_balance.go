package algorithms

import (
	// "fmt"
	"sort"
)

// len(curr) == len(deli) and k > 0
func DeliveryBalance(curr []int64, k int64) {
	var i, m int

	rs := make([]*int64, len(curr))
	for i = range curr {
		rs[i] = &(curr[i])
	}

	sort.Slice(rs, func(i, j int) bool { return *rs[i] < *rs[j] })

	i, m = 0, len(rs)-1

	for k > 0 {
		switch {
		case i < m && *rs[i] <= *rs[i+1]:
			*rs[i]++
			k--
		case i > 0 && *rs[i-1] <= *rs[i]:
			i--
		case i < m && *rs[i] > *rs[i+1]:
			i++
		case i == m && *rs[i-1] == *rs[i]:
			i = 0
			continue
		}
	}

	return
}
