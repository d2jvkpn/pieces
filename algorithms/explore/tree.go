package explore

import (
	"fmt"
)

type Node struct {
	V       int   // value
	P, L, R *Node // Parent, Left, Right
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
