package main

import (
	"fmt"
	"regexp"
)

func main() {
	x := regexp.MustCompile("(xyz)|(kk)")
	PrintResult(true, x.Match([]byte("xyz")))
	PrintResult(false, x.Match([]byte("kk")))
	PrintResult(false, x.Match([]byte("xk")))
	PrintResult(false, x.Match([]byte("xy")))

	// [0-9A-Za-z]{2}  (_[0-9]+)?  (.jpeg)|(.png)
	y := regexp.MustCompile("^[0-9A-Za-z]{2}(_[0-9]+)?(.jpeg)|(.png)$")
	PrintResult(true, y.Match([]byte("aa.png")))
	PrintResult(false, y.Match([]byte("aa_12.jpeg")))
	PrintResult(false, y.Match([]byte("aa_xx")))
}

var (
	m, n int
)

func PrintResult(newSect bool, x interface{}, a ...interface{}) {
	if newSect {
		m++
		n = 1
	} else {
		n++
	}

	str := fmt.Sprintf(">>> %d.%d %v\n", m, n, x)

	fmt.Printf(str, a...)
}
