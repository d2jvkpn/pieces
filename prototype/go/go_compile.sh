#! /usr/bin/env bash
set -eu -o pipefail

wd=$(pwd)

cat > tmp.go << EOF
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
EOF

go tool compile -S -N -l tmp.go
