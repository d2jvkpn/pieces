package explore

import (
// "fmt"
)

type TreeNode struct {
	Value  int    // represent status code
	Posi   [2]int // coordinate position
	Parent *TreeNode
	Left   *TreeNode
	Right  *TreeNode
}
