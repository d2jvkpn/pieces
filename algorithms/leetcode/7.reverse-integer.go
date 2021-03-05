package leetcode

import (
	"fmt"
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

func InstReverseNum() {
	fmt.Println(">>> InstReverseNum:")
	var num int
	num = 419

	fmt.Printf("    number = %d, out = %d\n", num, ReverseNum(num))
	num = -123
	fmt.Printf("    number = %d, out = %d\n", num, ReverseNum(num))
}
