package algorithms

import (
	"fmt"
)

type TreeNode struct {
	Value   int    // represent status code
	Posi    [2]int // coordinate position
	P, L, R *TreeNode
}
