package misc

import (
	"fmt"
	"testing"
)

func TestResData_t1(t *testing.T) {
	rd := NewResData(0, "OK")
	rd.SetData("name", "rover")

	rd.RequestId = "xxxx_1"
	rd.SetErrmsg(fmt.Errorf("something is wrong"))

	fmt.Println(rd.Pretty())
}
