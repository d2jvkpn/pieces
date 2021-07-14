#! /usr/bin/env bash
set -eu -o pipefail

wd=$(pwd)

echo """
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
	
	a := 1
	b := 2
	fmt.Println(a+b)
}
""" > tmp.go

go tool compile -S -N -l tmp.go
