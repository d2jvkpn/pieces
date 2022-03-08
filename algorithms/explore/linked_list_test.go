package explore

import (
	"fmt"
	"testing"
)

func TestLinkedList_a01(t *testing.T) {
	list := NewLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	fmt.Printf(">>> %#v\n    %#v\n", list, list.ToSlice())

	ok, value := list.Pop()
	fmt.Printf(">>> %t, %v\n    %#v\n", ok, value, list.ToSlice())
	
	list.Pop()
	list.Pop()
	fmt.Printf(">>> %#v\n    %#v\n", list, list.ToSlice())
}
