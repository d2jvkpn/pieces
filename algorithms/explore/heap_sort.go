package explore

import (
	"fmt"
)

func BuildTree(slice []int) (root *Node) {
	if len(slice) == 0 {
		return nil
	}

	var (
		n        int
		queue    []*Node
		bindNode func(*Node, *Node, *int)
		addSwap  func(*Node, *Node, bool)
	)

	queue = make([]*Node, len(slice))
	for i := range slice {
		queue[i] = NewNode(slice[i])
	}

	printQueue := func(queue []*Node) {
		ints := make([]int, len(queue))
		for i := range queue {
			ints[i] = queue[i].Value
		}
		fmt.Printf("    queue = %v\n", ints)
	}

	addSwap = func(node1, node2 *Node, greater bool) {
		if greater && node1.Value < node2.Value {
			fmt.Printf("    addSwap %d and %d\n", node1.Value, node2.Value)
			node1.Value, node2.Value = node2.Value, node1.Value
		}

		if node1.Parent != nil {
			addSwap(node1.Parent, node1, greater)
		}
	}

	bindNode = func(parent, node *Node, n *int) {
		switch {
		case parent.Left == nil:
			fmt.Printf("    setting %d.L = %d\n", parent.Value, node.Value)
			parent.Left, node.Parent = node, parent
		default:
			fmt.Printf("    setting %d.R = %d\n", parent.Value, node.Value)
			parent.Right, node.Parent = node, parent
			*n++
		}

		addSwap(parent, node, true)
	}

	root, n = queue[0], 0
	for _, v := range queue[1:] {
		bindNode(queue[n], v, &n)
	}

	printQueue(queue)

	return
}

func HeapSort(slice []int) (out []int) {
	if len(slice) < 2 {
		return slice
	}

	var (
		v       int
		popSwap func(*Node) (*Node, int)
	)

	root := BuildTree(slice)

	greater := func(node1, node2 *Node) bool {
		switch {
		case node1 == nil && node2 == nil:
			return false
		case node1 != nil && node2 == nil:
			return true
		case node1 == nil && node2 != nil:
			return false
		default:
			return node1.Value > node2.Value
		}
	}

	popSwap = func(node *Node) (out *Node, v int) {
		v = node.Value

		fmt.Printf("    popSwap node %v\n", node)

		switch {
		case node.Left == nil && node.Right == nil:
			if node.Parent != nil {
				if node.Parent.Left == node {
					node.Parent.Left = nil
				} else {
					node.Parent.Right = nil
				}
				fmt.Printf("    popSwap drop %d\n", node.Value)
			}
			out = nil
		case greater(node.Left, node.Right):
			fmt.Printf("    popSwap %d -> %d\n", node.Left.Value, node.Value)
			node.Value = node.Left.Value
			popSwap(node.Left)
			out = node
		default:
			fmt.Printf("    popSwap %d -> %d\n", node.Right.Value, node.Value)
			node.Value = node.Right.Value
			popSwap(node.Right)
			out = node
		}

		return
	}

	out = make([]int, 0, len(slice))
	for root != nil {
		root, v = popSwap(root)
		out = append(out, v)
		fmt.Printf("    append to out: %d, root = %s\n", v, root)
	}

	return out
}

func InstHeapSort1() {
	fmt.Println(">>> InstHeapSort1:")
	slice := []int{14, 33, 10, 27, 19, 35, 42, 44, 18, 17, 12, 28}
	fmt.Printf("    slice = %v\n", slice)

	out := HeapSort(slice)
	fmt.Printf("    out = %#v\n", out)
}
