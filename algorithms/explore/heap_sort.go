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

func (node *Node) NoC() (n int) {
	if node == nil {
		return 0
	}

	if node.L != nil {
		n++
	}

	if node.R != nil {
		n++
	}

	return
}

func AddNode(root, node *Node) {
	switch {
	case root.L == nil:
		fmt.Printf("    setting %d.L = %d\n", root.V, node.V)
		root.L, node.P = node, root

		return
	case root.R == nil:
		fmt.Printf("    setting %d.R = %d\n", root.V, node.V)
		root.R, node.P = node, root
		return
	}

	if root.L.NoC()-root.R.NoC() == 2 { // balance
		AddNode(root.R, node)
	} else {
		AddNode(root.L, node)
	}
}

func SwapNode(node1, node2 *Node) {
	fmt.Printf("    swap %d and %d\n", node1.V, node2.V)
	node1.V, node2.V = node2.V, node1.V
}

func HeapSort(slice []int) (out []int) {
	if len(slice) < 2 {
		return slice
	}

	var root *Node

	for i := range slice {
		if i == 0 {
			root = NewNode(slice[i])
			continue
		}

		AddNode(root, &Node{V: slice[i]})
	}

	fmt.Println(root)
	fmt.Println(root.L, root.R)
	fmt.Println(root.L.L, root.L.R, root.R.L, root.R.R)
	fmt.Println(root.L.L.L, root.L.L.R, root.L.R.L, root.L.R.R)

	return out
}

func InstHeapSort() {
	fmt.Println(">>> InstQuickSort:")
	slice := []int{14, 33, 10, 27, 19, 35, 42, 44, 18}
	fmt.Printf("    slice = %v\n", slice)

	out := HeapSort(slice)
	fmt.Printf("    out = %#v\n", out)
}
