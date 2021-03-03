package main

import (
	"fmt"
	"os"
	"testing"
)

func TestLoadBashrc(t *testing.T) {
	var err error
	if err = LoadBashrc("test_data/bashrc"); err != nil {
		t.Fatal(err)
	}

	fmt.Println("A:", os.Getenv("A"))
	fmt.Println("B:", os.Getenv("B"))
}
