package explore

import (
	"fmt"
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

func NewLinkedListNode(value int64) *LinkedListNode {
	return &LinkedListNode{Value: value}
}

func NewLinkedList() LinkedList {
	return LinkedList{}
}

func (list LinkedList) String() string {
	head, tail := list.Head, list.Tail
	if head == nil {
		return fmt.Sprintf("head=0, tail=0, len=0")
	}
	return fmt.Sprintf("head=%d, tail=%d, len=%d", head.Value, tail.Value, list.Len)
}

func (list *LinkedList) Append(value int64) *LinkedList {
	node := NewLinkedListNode(value)

	if list.Head == nil && list.Tail == nil {
		list.Head = node
	} else if list.Head.Next == nil && list.Tail != nil {
		list.Head.Next, node.Prev = node, list.Head
	} else {
		list.Tail.Next, node.Prev = node, list.Tail
	}

	list.Tail = node
	list.Len++
	return list
}

func (list *LinkedList) Pop() (value int64, ok bool) {
	if list.Tail == nil {
		return 0, false
	}

	value, ok = list.Tail.Value, true
	if list.Head.Next == nil {
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

	for node := list.Head; node != nil; node = node.Next {
		slice = append(slice, node.Value)
	}

	return
}

func (list *LinkedList) Index(n int) (value int64, ok bool) {
	if n > list.Len-1 {
		return 0, false
	}

	for i, node := 0, list.Head; i < list.Len; i, node = i+1, node.Next {
		if i == n {
			return node.Value, true
		}
	}

	return
}

func (list *LinkedList) Drop(n int) (value int64, ok bool) {
	if n > list.Len-1 {
		return 0, false
	}

	for i, node := 0, list.Head; i < list.Len; i, node = i+1, node.Next {
		if i == n {
			if prev, next := node.Prev, node.Next; prev == next {
				list.Head, list.Tail, list.Len = nil, nil, 0
			} else {
				prev.Next, next.Prev, list.Len = next, prev, list.Len-1
			}
			return node.Value, true
		}
	}

	return
}
