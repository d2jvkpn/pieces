package misc

import (
	"fmt"
	"testing"
)

func TestResData_t1(t *testing.T) {
	rd := NewResData(0, "OK")
	rd.Set("name", "rover")

	rd.RequestId = "xxxx_1"
	rd.Err = fmt.Errorf("something is wrong")

	fmt.Printf(">>> TestResData_t1:\n%s\n", rd)
	fmt.Printf(">>> TestResData_t1:\n%s\n", rd.Pretty())
}

func TestResData_t2(t *testing.T) {
	var rd *ResData
	var err error = rd

	fmt.Printf(">>> TestResData_t2:\n%v\n", err)
}
