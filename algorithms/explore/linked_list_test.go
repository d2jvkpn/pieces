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

	value, ok := list.Pop()
	fmt.Printf(">>> %t, %v\n    %#v\n", ok, value, list.ToSlice())

	list.Pop()
	list.Pop()
	fmt.Printf(">>> %#v\n    %#v\n", list, list.ToSlice())
}

func TestLinkedList_a02(t *testing.T) {
	list := NewLinkedList()
	list.Append(1).Append(2).Append(3).Append(4).Append(5).Append(6).Append(7).Append(8)

	value, ok := list.Index(3)
	fmt.Printf(">>> %t, %v\n    %#v\n", ok, value, list.ToSlice())
}
