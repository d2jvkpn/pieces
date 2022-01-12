package main

import (
	"fmt"
)

type String string

func (str *String) Length() int {
	return len(*str)
}

func main() {
	var str String = "Hello, world!"
	fmt.Println(str.Length())
}
