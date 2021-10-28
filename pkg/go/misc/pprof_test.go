package misc

import (
	// "fmt"
	"testing"
)

func TestPprof(t *testing.T) {
	pp := NewPprof(":1030")

	t.Fatal(pp.Run())
}
