package explore

import (
	"fmt"
)

func BuildTree2(slice []int, asc bool) (root *Node) {
	if len(slice) == 0 {
		return nil
	}

	var (
		cursor   *int
		queue    []*Node
		bindNode func(*Node, *Node, *int)
		pushSwap func(*Node, *Node, bool)
	)

	queue = make([]*Node, len(slice))
	for i := range slice {
		queue[i] = NewNode(slice[i])
	}

	printQueue := func(queue []*Node) {
		ints := make([]int, len(queue))
		for i := range queue {
			ints[i] = queue[i].V
		}
		fmt.Printf("    queue = %v\n", ints)
	}

	pushSwap = func(node1, node2 *Node, lt bool) { // lt => less than
		if (lt && node1.V > node2.V) || (!lt && node1.V < node2.V) {
			fmt.Printf("    pushSwap node: %d with %d\n", node1.V, node2.V)
			node1.V, node2.V = node2.V, node1.V
		}

		if node1.P != nil {
			pushSwap(node1.P, node1, lt)
		}
	}

	bindNode = func(parent, node *Node, cursor *int) {
		switch {
		case parent.L == nil:
			fmt.Printf("    bindNode: %d.L = %d\n", parent.V, node.V)
			parent.L, node.P = node, parent
		default:
			fmt.Printf("    bindNode: %d.R = %d\n", parent.V, node.V)
			parent.R, node.P = node, parent
			*cursor++
		}

		pushSwap(parent, node, asc)
	}

	root, cursor = queue[0], new(int)
	for _, v := range queue[1:] {
		bindNode(queue[*cursor], v, cursor)
	}

	printQueue(queue)
	return
}

func HeapSort2(slice []int, asc bool) (out []int) {
	if len(slice) < 2 {
		return slice
	}

	var (
		v       int
		root    *Node
		popSwap func(*Node) (*Node, int)
	)

	root = BuildTree2(slice, asc)

	choose := func(node1, node2 *Node, lt bool) (out *Node) {
		switch {
		case node1 == nil && node2 == nil:
			return nil
		case node1 != nil && node2 == nil:
			return node1
		case node1 == nil && node2 != nil:
			return node2
		case (node1.V < node2.V && lt) || (node1.V > node2.V && !lt): //!!!
			return node1
		default:
			return node2
		}
	}

	dropNode := func(node *Node) {
		if node.P == nil { // root nodes
			return
		}

		if node.P.L == node {
			node.P.L = nil
		} else {
			node.P.R = nil
		}
		fmt.Printf("    drop node: %s\n", node)
	}

	popSwap = func(node *Node) (out *Node, v int) {
		v = node.V
		if node.L == nil && node.R == nil {
			dropNode(node)
			return nil, v
		}

		if x := choose(node.L, node.R, asc); x != nil {
			fmt.Printf("    popSwap 2 nodes: %s with %s\n", node, x)
			node.V = x.V
			popSwap(x)
		}

		out = node
		return
	}

	out = make([]int, 0, len(slice))
	for root != nil {
		root, v = popSwap(root)
		out = append(out, v)
		fmt.Printf("    append to out: %d, root = %s\n", v, root)
	}

	return
}

func InstHeapSort2() {
	fmt.Println(">>> InstHeapSort2:")
	slice := []int{14, 33, 10, 27, 19, 35, 42, 44, 18, 17, 12, 28}
	fmt.Printf("    slice = %v\n", slice)

	out := HeapSort2(slice, true)
	fmt.Printf("    out = %#v\n", out)
}

func InstHeapSort3() {
	fmt.Println(">>> InstHeapSort3:")
	slice := []int{14, 33, 10, 27, 19, 35, 42, 44, 18, 17, 12, 28}
	fmt.Printf("    slice = %v\n", slice)

	out := HeapSort2(slice, false)
	fmt.Printf("    out = %#v\n", out)
}
