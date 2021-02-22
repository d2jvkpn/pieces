package leetcode

import (
	"math"
)

func ReverseNum(x int) (out int) {
	num, isNeg := x, x < 0
	if isNeg {
		num = -num
	}

	arr := make([]int, 0)
	for {
		arr = append(arr, num%10)
		if num /= 10; num == 0 {
			break
		}
	}

	// fmt.Println(arr)
	for i := range arr {
		out += arr[i] * int(math.Pow10(len(arr)-i-1))
	}

	if isNeg {
		out = -out
	}

	return
}
