package misc

import (
	"fmt"
	"testing"
)

func TestVector(t *testing.T) {
	list := Vector[string]{"a", "b", "c"}

	fmt.Printf("%v, %v\n", *list.First(), *list.Last())

	var p *int
	p2 := &p

	fmt.Println(p == nil, p2 == nil)
}
