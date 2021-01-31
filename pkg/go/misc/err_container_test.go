package misc

import (
	"context"
	"fmt"
	"testing"
)

type xxxx struct {
	context.Context
}

func xxxxWithErrContainer(ctx context.Context) (x *xxxx) {
	type C struct {
		Username string
		Password string
	}

	c := &C{Username: "rover", Password: "123456"}

	return &xxxx{Context: context.WithValue(
		ctx,
		"Error",
		NewErrContainer(c, "simple password").SetLevel(2),
	)}
}

func TestErrContainer(t *testing.T) {
	x := xxxxWithErrContainer(context.TODO())
	err, _ := x.Value("Error").(interface {
		Value() interface{}
		Error() string
		Level() int
	})

	if err == nil {
		fmt.Println("Error is nil.")
	} else {
		fmt.Printf(">>> Error: %q, Level: %d\n", err.Error(), err.Level())
		fmt.Printf("    Value: %#v\n", err.Value())
	}
}
