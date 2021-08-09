package main

import (
	"fmt"
)

func main() {
	d := Data{}
	d1 := &d
	d2 := d1

	d1 = nil
	PrintData(d1) // <nil>
	PrintData(d2) // &{0}
}

type Data struct {
	V int64
}

func PrintData(p *Data) {
	fmt.Println(p)
}
