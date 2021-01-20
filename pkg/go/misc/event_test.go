package misc

import (
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	err := NewError(fmt.Errorf("something is wrong!"))

	PrintJSON(err)
}

func TestEvent(t *testing.T) {
	event := NewEvent(
		"find user", 1,
		WithError(fmt.Errorf("not found")),
		WithData(map[string]int64{"user_id": 20201230}),
	)

	PrintJSON(event)
}
