package misc

import (
	"fmt"
	"testing"
)

func TestProjectDir(t *testing.T) {
	p, err := ProjectDir()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(p)
}
