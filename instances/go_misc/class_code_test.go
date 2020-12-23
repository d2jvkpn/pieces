package rover

import (
	"fmt"
	"testing"
)

func TestClassCode_x1(t *testing.T) {
	ec, err := NewClassCode(ClassElem{"A", true}, ClassElem{"B", false},
		ClassElem{"C", true}, ClassElem{"D", true})

	fmt.Println(err)
	fmt.Printf("%+v\n", ec)
	fmt.Printf("%+v\n", ec.Elems())
	fmt.Printf("%+v\n", ec.Bools())

	fmt.Println(ec.Value("A"))
	fmt.Println(ec.Value("B"))
	fmt.Println(ec.Value("X"))
}

func TestClassCode_x2(t *testing.T) {
	ec, _ := NewClassCode(ClassElem{"A", true}, ClassElem{"B", false},
		ClassElem{"C", true}, ClassElem{"D", true})

	fmt.Printf("%+v\n", ec.Bools())

	ec.ClearAt(1)
	fmt.Printf("%+v\n", ec.Bools())

	ec.ToggleAt(1)
	fmt.Printf("%+v\n", ec.Bools())

	ec.ToggleAt(3)
	fmt.Printf("%+v\n", ec.Bools())
}
