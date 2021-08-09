package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

func main() {
	var (
		bts            []byte
		e1, e2, e3, e4 error
	)

	e2 = errors.New("something wrong")
	e3 = fmt.Errorf("Something Wrong")

	bts, _ = json.Marshal(e1)
	fmt.Printf("%T, %s\n", e1, bts)

	bts, _ = json.Marshal(e2)
	fmt.Printf("%T, %s\n", e2, bts) // *errors.errorString, {}

	bts, _ = json.Marshal(e3)
	fmt.Printf("%T, %s\n", e3, bts) // *errors.errorString, {}

	d1 := &D{99}
	e4 = d1
	bts, _ = json.Marshal(e4)
	fmt.Printf("%T, %s\n", e4, bts) // *main.D, {"Value":99}

	e4 = fmt.Errorf(">>> error: %w", d1)
	bts, _ = json.Marshal(e4)
	fmt.Printf("%T, %s\n", e4, bts) // *fmt.wrapError, {}
}

type D struct {
	Value int64
}

func (d *D) Error() string {
	if d.Value == 0 {
		return "<nil>"
	}

	return fmt.Sprintf("current value: %d", d.Value)
}
