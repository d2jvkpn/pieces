package explore

import (
	"fmt"
)

type Node struct {
	Value  int64
	Color  bool // false represents black
	Parent *Node
	Left   *Node
	Right  *Node
}

/* Rules:
1. Every node has a colour either red or black.
2. The root of the tree is always black.
3. There are no two adjacent red nodes (A red node cannot have a red parent or red child).
4. Every path from a node (including root) to any of its descendants NULL nodes has the same
number of black nodes.
5. All leaf nodes are black nodes.
*/
func (node *Node) Valid() (err error) {
	if node.Parent == nil && !node.Parent.Color {
		return fmt.Errorf("rule 2")
	}

	if node.Color {
		err = fmt.Errorf("rule 3")
		switch {
		case node.Parent != nil && node.Parent.Color:
			return err
		case node.Left != nil && node.Left.Color:
			return err
		case node.Right != nil && node.Right.Color:
			return err
		}
		err = nil
	}

	// ?rule 4
	return nil
}
