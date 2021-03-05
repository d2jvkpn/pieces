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
		n            int
		queue        []*Node
		bindNode     func(*Node, *Node, *int)
		heapSortSwap func(*Node, *Node, bool)
	)

	queue = make([]*Node, len(slice))
	for i := range slice {
		queue[i] = NewNode(slice[i])
	}

	printQueue := func(queue []*Node) {
		strs := make([]int, len(queue))
		for i := range queue {
			strs[i] = queue[i].V
		}
		fmt.Printf("    queue = %v\n", strs)
	}

	heapSortSwap = func(node1, node2 *Node, greater bool) {
		if greater && node1.V < node2.V {
			fmt.Printf("    swap %d and %d\n", node1.V, node2.V)
			node1.V, node2.V = node2.V, node1.V
		}

		if node1.P != nil {
			heapSortSwap(node1.P, node1, greater)
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

		heapSortSwap(parent, node, true)
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

	root := BuildTree(slice)
	fmt.Println(root)

	fmt.Println(root.L, root.R)
	fmt.Println(root.L.L, root.L.R, root.R.L, root.R.R)
	fmt.Println(root.L.L.L, root.L.L.R, root.L.R.L, root.L.R.R)

	return out
}

func InstHeapSort() {
	fmt.Println(">>> InstQuickSort:")
	slice := []int{14, 33, 10, 27, 19, 35, 42, 44, 18, 17, 12, 28}
	fmt.Printf("    slice = %v\n", slice)

	out := HeapSort(slice)
	fmt.Printf("    out = %#v\n", out)
}
