package rover

import (
	"fmt"
	"testing"
	"time"
)

func TestUpdateTrigger(t *testing.T) {
	fn := func() error {
		fmt.Println("Hello", time.Now())
		return nil
	}

	ut, _ := NewUpdateTrigger(fn, 10*time.Second, true, true)
	fmt.Println(">>> 1")
	fmt.Println(ut.Update())

	fmt.Println(">>> 2")
	fmt.Println(ut.Update())

	time.Sleep(10 * time.Second)
	fmt.Println(">>> 3")
	fmt.Println(ut.Update())

	time.Sleep(12 * time.Second)
	fmt.Println(">>> 4")
	fmt.Println(ut.Update())
}
