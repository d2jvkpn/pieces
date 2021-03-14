package explore

import (
	"fmt"
)

type Node struct {
	Value       int
	Parent      *Node
	Left, Right *Node
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}

func NewNode2(value int, parent *Node, do func(node *Node)) (out *Node) {
	out = &Node{Value: value, Parent: parent}
	if do == nil {
		return
	}

	for ; parent != nil; parent = parent.Parent {
		do(parent)
	}
	return
}

func (node *Node) String() string {
	if node == nil {
		return "."
	}
	return fmt.Sprintf("%d", node.Value)
}
