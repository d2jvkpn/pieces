package main

import (
	"fmt"
	"runtime"
)

func foo(n int) {
	fn, file, no, ok := runtime.Caller(n)
	details := runtime.FuncForPC(fn)

	if ok {
		fmt.Printf("called from %s: %s[%d]\n", file, details.Name(), no)
	} else {
		foo(n - 1)
	}
}

func foo2(n int) {
	foo(n)
}

func main() {
	foo2(1)
	foo2(2)
	foo2(10)
}
