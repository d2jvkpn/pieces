package interview

import (
	"fmt"
)

type Node struct {
	Value int      // 0: be cleared, 1: origin value
	Posi  [2]int   // position
	URDL  [4]*Node // up, right, down, left
}

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
			v.Clear(node)
			v.Value = 0
		}
	}
}

func MatrixNeighborOnes(matrix [][]int) (num int) {
	nodes := make(map[[2]int]*Node)

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				continue
			}
			posi := [2]int{i, j}
			nodes[posi] = &Node{Value: 1, Posi: posi}
		}
	}

	for k := range nodes {
		i, j := k[0], k[1]
		nodes[k].URDL[0] = nodes[[2]int{i - 1, j}]
		nodes[k].URDL[1] = nodes[[2]int{i, j + 1}]
		nodes[k].URDL[2] = nodes[[2]int{i + 1, j}]
		nodes[k].URDL[3] = nodes[[2]int{i, j - 1}]
	}

	for k := range nodes {
		fmt.Printf(">>> Iterating at: %v\n", k)
		nodes[k].Clear(nil)

		if nodes[k].Value == 1 {
			num++
		}
	}

	return
}

func InstNeighborOnes() {
	// expect InstNeighborOnes(matrix) = 4
	matrix := [][]int{
		{1, 0, 0, 0, 1},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 1, 0},
	}

	fmt.Println(MatrixNeighborOnes(matrix))
}
