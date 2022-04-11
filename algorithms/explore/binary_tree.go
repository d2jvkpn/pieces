package explore

import (
	"fmt"
)

type Node struct {
	Value  int64
	Parent *Node
	Left   *Node
	Right  *Node
}

func NewNode(value int64) *Node {
	return &Node{Value: value}
}

func (node *Node) IsRoot() bool {
	return node.Parent == nil
}

func (node *Node) Add(value int64) *Node {
	if value <= node.Value {
		if node.Left == nil {
			node.Left = NewNode(value)
			node.Left.Parent = node
		} else {
			node.Left.Add(value)
		}
	} else if node.Right == nil {
		node.Right = NewNode(value)
		node.Right.Parent = node
	} else {
		node.Right.Add(value)
	}

	return node
}

func (node *Node) String() (out string) {
	if node == nil {
		return "nil"
	}

	out = fmt.Sprintf("%d", node.Value)
	if node.Parent != nil {
		out = fmt.Sprintf("%s(%d)", out, node.Parent.Value)
	}

	if node.Left == nil {
		out = "nil-" + out
	} else {
		out = fmt.Sprintf("%d-%s", node.Left.Value, out)
	}

	if node.Right == nil {
		out = out + "-nil"
	} else {
		out = fmt.Sprintf("%s-%d", out, node.Right.Value)
	}

	return
}

func (node *Node) Search(value int64) (out *Node, n int) {
	out = node.search(value, &n)
	return
}

func (node *Node) search(value int64, np *int) (out *Node) {
	if node == nil {
		return nil
	}

	switch {
	case node.Value == value:
		return node
	case node.Value > value:
		*np++
		return node.Left.search(value, np)
	default:
		*np++
		return node.Right.search(value, np)
	}
}

func main() {
	root := NewNode(8).Add(1).Add(9).Add(3).Add(10).Add(15).Add(5)

	fmt.Println(root)
	fmt.Println(root.Left, root.Left.Right, root.Left.Right.Right)
	fmt.Println(root.Right, root.Right.Right)

	fmt.Println(">>> search 5")
	fmt.Println(root.Search(5))

	fmt.Println(">>> search 13")
	fmt.Println(root.Search(13))
}
