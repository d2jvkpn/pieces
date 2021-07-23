package errorx

import (
	"fmt"
	"testing"
)

func TestInvalidParameter(t *testing.T) {
	errx := InvalidParameter(fmt.Errorf("pageSize is invalid"), "xxxx")

	fmt.Printf("%+v\n", errx)
}
