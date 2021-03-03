package main

import (
	"fmt"
	"testing"
)

func TestPrintErr(t *testing.T) {
	PrintErr("Hello", nil, 1, 2, 3)
	err := fmt.Errorf("???")
	PrintErr("Hello", err, 1, 2, 3)
	PrintErr("Hello", err)
}

func TestLogErr(t *testing.T) {
	LogErr("Hello", nil)
	err := fmt.Errorf("???")
	LogErr("Hello", err)
}
