package misc

import (
	"fmt"
	"testing"
)

func _fn2() {
	defer GetPanic()
	_fn1()
}

func _fn1() {
	var mySlice []int
	j := mySlice[0]

	fmt.Printf("Hello, playground %d", j)
}

// go test -run  TestGetPanic | sed '$d' | sed '$d' | jq .panicStack | xargs -i printf {}
func TestGetPanic(t *testing.T) {
	_fn2()
}
