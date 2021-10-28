package misc

import (
	// "fmt"
	"testing"
)

func TestPprof(t *testing.T) {
	pp := NewPprof(":5060")

	pp.Run()
	t.Fatal(pp.Err())
}
