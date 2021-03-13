package explore

import (
	"fmt"
)

type Node struct {
	Value  int
	Parent *Node
	Left   *Node
	Right  *Node
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}

func (node *Node) String() string {
	if node == nil {
		return "."
	}
	return fmt.Sprintf("%d", node.Value)
}
