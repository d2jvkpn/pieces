package main

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func foo(n int) {
	fn, file, line, ok := runtime.Caller(n)
	details := runtime.FuncForPC(fn)

	if ok {
		fmt.Printf("called from \"%s (%s [%d])\"\n", details.Name(), filepath.Base(file), line)
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
