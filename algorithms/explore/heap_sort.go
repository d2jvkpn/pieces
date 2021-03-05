package explore

import (
	"fmt"
)

type Node struct {
	V       int
	P, L, R *Node
}

func NewNode(value int) *Node {
	return &Node{V: value}
}

func (node *Node) String() string {
	if node == nil {
		return "."
	}
	return fmt.Sprintf("%d", node.V)
}

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
			ints[i] = queue[i].V
		}
		fmt.Printf("    queue = %v\n", ints)
	}

	addSwap = func(node1, node2 *Node, greater bool) {
		if greater && node1.V < node2.V {
			fmt.Printf("    addSwap %d and %d\n", node1.V, node2.V)
			node1.V, node2.V = node2.V, node1.V
		}

		if node1.P != nil {
			addSwap(node1.P, node1, greater)
		}
	}

	bindNode = func(parent, node *Node, n *int) {
		switch {
		case parent.L == nil:
			fmt.Printf("    setting %d.L = %d\n", parent.V, node.V)
			parent.L, node.P = node, parent
		default:
			fmt.Printf("    setting %d.R = %d\n", parent.V, node.V)
			parent.R, node.P = node, parent
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
			return node1.V > node2.V
		}
	}

	popSwap = func(node *Node) (out *Node, v int) {
		v = node.V

		fmt.Printf("    popSwap node %v\n", node)

		switch {
		case node.L == nil && node.R == nil:
			if node.P != nil {
				if node.P.L == node {
					node.P.L = nil
				} else {
					node.P.R = nil
				}
				fmt.Printf("    popSwap drop %d\n", node.V)
			}
			out = nil
		case greater(node.L, node.R):
			fmt.Printf("    popSwap %d -> %d\n", node.L.V, node.V)
			node.V = node.L.V
			popSwap(node.L)
			out = node
		default:
			fmt.Printf("    popSwap %d -> %d\n", node.R.V, node.V)
			node.V = node.R.V
			popSwap(node.R)
			out = node
		}

		return
	}

	out = make([]int, 0, len(slice))
	for root != nil {
		root, v = popSwap(root)
		out = append(out, v)
		fmt.Printf("    ~~~ value = %d, root = %s\n", v, root)
	}

	return out
}

func InstHeapSort() {
	fmt.Println(">>> InstQuickSort:")
	slice := []int{14, 33, 10, 27, 19, 35, 42, 44, 18, 17, 12, 28}
	fmt.Printf("    slice = %v\n", slice)

	out := HeapSort(slice)
	fmt.Printf("    out = %#v\n", out)
}
