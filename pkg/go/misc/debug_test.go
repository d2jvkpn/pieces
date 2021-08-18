package misc

import (
	"fmt"
	"testing"
)

func _fn2(n int) {
	defer GetPanic(n)
	_fn1()
}

func _fn1() {
	var mySlice []int
	j := mySlice[0]

	fmt.Printf("Hello, playground %d", j)
}

func wrapError2(err error) error {
	return WrapError(err)
}

func TestWrapError_t1(t *testing.T) {
	err := WrapError(fmt.Errorf("something is wrong"))
	fmt.Println(err)
}

func TestWrapError_t2(t *testing.T) {
	err := wrapError2(fmt.Errorf("something is wrong 2"))
	fmt.Println(err)
}

// go test -run  TestGetPanic_t1 | sed '$d' | sed '$d' | jq .panicStack | xargs -i printf {}
func TestGetPanic_t1(t *testing.T) {
	_fn2(1)
}

func TestGetPanic_t2(t *testing.T) {
	_fn2(2)
}

func TestGetPanic_t3(t *testing.T) {
	_fn2(3)
}

func TestGetPanic_t4(t *testing.T) {
	_fn2(20)
}
