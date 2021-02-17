package main

import (
	_ "embed"
	"fmt"
)

//go:embed fs/hello.txt
var bts []byte

func main() {
	fmt.Print(string(bts))
}
