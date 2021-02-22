package leetcode

import (
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	toInt := func(node *ListNode) (value int) {
		value = node.Val

		next := node.Next
		for i := 1; next != nil; i++ {
			value += next.Val * int((math.Pow(10, float64(i))))
			next = next.Next
		}

		return
	}

	toNode := func(value int) (out *ListNode) {
		arr := make([]*ListNode, 0)
		for {
			arr = append(arr, &ListNode{Val: value % 10})
			if value /= 10; value == 0 {
				break
			}
		}

		for i := 0; i < len(arr)-1; i++ {
			arr[i].Next = arr[i+1]
		}

		return arr[0]
	}

	return toNode(toInt(l1) + toInt(l2))
}
