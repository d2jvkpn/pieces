package explore

import (
// "fmt"
)

type LinkedListNode struct {
	Value int64
	Prev  *LinkedListNode
	Next  *LinkedListNode
}

type LinkedList struct {
	Head *LinkedListNode
	Tail *LinkedListNode
	Len  int
}

func NewLinkedListNode(value int64) LinkedListNode {
	return LinkedListNode{Value: value}
}

func NewLinkedList() LinkedList {
	return LinkedList{}
}

func (list *LinkedList) Append(value int64) *LinkedList {
	node := NewLinkedListNode(value)
	// fmt.Println(node)

	if list.Tail != nil {
		list.Tail.Next, node.Prev = &node, list.Tail
	}
	list.Tail = &node

	if list.Head == nil {
		list.Head = &node
	}
	list.Len++
	return list
}

func (list *LinkedList) Pop() (value int64, ok bool) {
	if list.Tail == nil {
		return 0, false
	}

	ok, value = true, list.Tail.Value
	if list.Head == list.Tail {
		list.Head, list.Tail, list.Len = nil, nil, 0
		return
	}

	list.Tail = list.Tail.Prev
	list.Tail.Next = nil
	list.Len--
	return
}

func (list *LinkedList) ToSlice() (slice []int64) {
	slice = make([]int64, 0, list.Len)

	for node := list.Head; node != nil && node != node.Next; node = node.Next {
		// fmt.Println("~~~", node.Value)
		slice = append(slice, node.Value)

	}
	return
}

func (list *LinkedList) Index(n int) (value int64, ok bool) {
	if n > list.Len-1 {
		return 0, false
	}

	for i, node := 0, list.Head; i < list.Len; i++ {
		if i == n {
			return node.Value, true
		}
		node = node.Next
	}

	return
}
