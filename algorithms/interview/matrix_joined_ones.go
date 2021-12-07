package interview

import (
	"fmt"
)

type Node struct {
	Value int      // 1 -> origin value, 0 -> be cleared
	Posi  [2]int   // position
	URDL  [4]*Node // up, right, down, left
}

type NodesMatrix struct {
	n, m  int
	nodes map[[2]int]*Node
}

// clear neighbors
func (node *Node) Clear(previous *Node) {
	fmt.Printf("    clearing: Node -> %v, Previous -> %v\n", node.Posi, previous)
	if node.Value == 0 {
		return
	}

	for _, v := range node.URDL {
		if v == nil || v.Value == 0 || v == previous {
			continue
		}

		if v.Value == 1 {
			v.Clear(node) // recurssion
			v.Value = 0
		}
	}
}

func (nm *NodesMatrix) Rows() int {
	return nm.n
}

func (nm *NodesMatrix) Columns() int {
	return nm.m
}

// create NodesMatrix from two dimensional array
func NewNodesMatrix(matrix [][]int) (nm *NodesMatrix) {
	nm = &NodesMatrix{
		n:     len(matrix),    // rows
		m:     len(matrix[0]), // columns
		nodes: make(map[[2]int]*Node),
	}

	// create nm
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				continue
			}
			posi := [2]int{i, j}
			nm.nodes[posi] = &Node{Value: 1, Posi: posi}
		}
	}

	// fill neighbors
	nodes := nm.nodes
	for k := range nodes {
		i, j := k[0], k[1]
		nodes[k].URDL[0] = nodes[[2]int{i - 1, j}] // up, may be nil
		nodes[k].URDL[1] = nodes[[2]int{i, j + 1}] // right
		nodes[k].URDL[2] = nodes[[2]int{i + 1, j}] // down
		nodes[k].URDL[3] = nodes[[2]int{i, j - 1}] // left
	}

	return
}

// clear neighbors and count nodes which .Value  == 1
func (nm *NodesMatrix) ClearNeighbors() (num int) {
	nodes := nm.nodes
	for k := range nodes {
		fmt.Printf(">>> Iterating at: %v\n", k)
		nodes[k].Clear(nil)

		if nodes[k].Value == 1 {
			num++
		}
	}

	return
}

// instance
func InstMatrixJoinedOnes() {
	fmt.Println(">>> InstMatrixJoinedOnes:")
	// expect output: 4
	matrix := [][]int{
		{1, 0, 0, 0, 1},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{0, 0, 0, 1, 0},
	}
	fmt.Printf("    matrix = %v\n", matrix)

	nm := NewNodesMatrix(matrix)
	num := nm.ClearNeighbors()
	fmt.Printf("    number = %d\n", num)
}
