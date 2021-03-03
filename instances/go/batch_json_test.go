package main

import (
	"fmt"
	"testing"
)

func TestBatchJSON(t *testing.T) {
	x := BatchJSON(6, func(bts []byte) error {
		fmt.Println(">>> batch", string(bts))
		return nil
	})

	for i := 0; i < 101; i++ {
		x(fmt.Sprintf("%d", i+1), true)
	}

	x("", false)
}
