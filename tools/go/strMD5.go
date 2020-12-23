package main

import (
	"crypto/md5"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "Please provide a string!\n")
		os.Exit(2)
	}

	fmt.Printf("%X\n", md5.Sum([]byte(os.Args[1])))
}
