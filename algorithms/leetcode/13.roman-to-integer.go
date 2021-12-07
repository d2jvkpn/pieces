package leetcode

import (
	"fmt"
)

func RomanToInteger(str string) (num int) {
	mp := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50,
		'C': 100, 'D': 500, 'M': 1000,
	}

	for i := 0; i < len(str); i++ {
		v, n := mp[str[i]], 0
		if i < len(str)-1 {
			n = mp[str[i+1]]
		}

		if (v == 1 || v%10 == 0) && (v*5 == n || v*10 == n) {
			v = n - v
			i++
		}

		num += v
	}

	return
}

func InstRomanToInteger() {
	fmt.Println(">>> InstRomanToInteger:")

	for _, str := range []string{"III", "IX", "LVIII"} {
		fmt.Printf("    roman = %s, number = %d\n", str, RomanToInteger(str))
	}
}
