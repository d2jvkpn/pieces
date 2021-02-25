package interview

import (
	"fmt"
)

type Node struct {
	Value int      // 0: be cleared, 1: origin value
	Posi  [2]int   // position
	URDL  [4]*Node // up, right, down, left
}

type NodesMatrix map[[2]int]*Node

// clear neighbors
func (node *Node) Clear(parent *Node) {
	fmt.Printf("    clearing: Node -> %v, Parent -> %v\n", node.Posi, parent)
	if node.Value == 0 {
		return
	}

	for _, v := range node.URDL {
		if v == nil || v.Value == 0 || v == parent {
			continue
		}

		if v.Value == 1 {
			v.Clear(node) // recurssion
			v.Value = 0
		}
	}
}

// create NodesMatrix from two dimensional array
func NewNodesMatrix(matrix [][]int) (nm NodesMatrix) {
	nm = make(map[[2]int]*Node)

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				continue
			}
			posi := [2]int{i, j}
			nm[posi] = &Node{Value: 1, Posi: posi}
		}
	}

	return
}

// clear neighbors and count nodes which .Value  == 1
func (nm NodesMatrix) ClearNeighbors() (num int) {
	for k := range nm {
		i, j := k[0], k[1]
		nm[k].URDL[0] = nm[[2]int{i - 1, j}]
		nm[k].URDL[1] = nm[[2]int{i, j + 1}]
		nm[k].URDL[2] = nm[[2]int{i + 1, j}]
		nm[k].URDL[3] = nm[[2]int{i, j - 1}]
	}

	for k := range nm {
		fmt.Printf(">>> Iterating at: %v\n", k)
		nm[k].Clear(nil)

		if nm[k].Value == 1 {
			num++
		}
	}

	return
}

// instance
func InstNeighborOnes() {
	// expect InstNeighborOnes(matrix) = 4
	matrix := [][]int{
		{1, 0, 0, 0, 1},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 1, 0},
	}

	nm := NewNodesMatrix(matrix)
	fmt.Println(nm.ClearNeighbors())
}
