package main

import (
	"fmt"
	"testing"
)

func TestNewCMSE(t *testing.T) {
	cmse := NewCMSE()
	fmt.Println(cmse.String())

	cmse.SetD("a", 1213)
	fmt.Println(cmse.String())

	cmse.E = fmt.Errorf("something wrong")
	fmt.Println(cmse.String())
}
