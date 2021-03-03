package main

import (
	"fmt"
	"math"
	"strings"
)

// ASCII code
var base62Codes = []int{
	48, 49, 50, 51, 52, 53, 54, 55, 56, 57, // [0-9]
	65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, // [A-Z]
	81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
	97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, // [a-z]
	110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122,
}

func Base62Encode(num int64) (result string) {
	var (
		rem int64
		bts []byte
		lt0 bool // less than 0
	)

	if num < 0 {
		lt0 = true
		num = -num
	}

	bts = make([]byte, 0, 10)
	for {
		num, rem = num/62, num%62
		bts = append(bts, byte(base62Codes[rem]))
		if num == 0 {
			break
		}
	}

	for i, j := 0, len(bts)-1; i < j; i, j = i+1, j-1 {
		bts[i], bts[j] = bts[j], bts[i]
	}

	result = string(bts)
	if lt0 {
		result = "-" + result
	}

	return
}

func Base62Decode(str string) (num int64, err error) {
	var (
		i, p int
		bts  []byte
		lt0  bool // less than 0
	)

	bts = []byte(str)
	if strings.HasPrefix(str, "-") {
		lt0 = true
		bts = bts[1:]
	}

	for i = range bts {
		if p = IndexInt(base62Codes, int(bts[i])); p < 0 {
			err = fmt.Errorf("invalid character to decode: %q", bts[i])
			return
		}

		num += int64(p) * int64(math.Pow(float64(62), float64(len(bts)-1-i)))
	}

	if lt0 {
		num = -num
	}

	return
}
