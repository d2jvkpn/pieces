package misc

import (
	"fmt"
	"testing"
)

func TestNewError2(t *testing.T) {
	err := NewError2(fmt.Errorf("something is wrong"))

	fmt.Println(err)
}
