package explore

import (
	"fmt"
)

// Brеаdth First Trаvеrsаl
func (node *GNode) BFS() {
	queue := []*GNode{node}

	// push one node and set Status = 1
	visit := func(n1 *GNode) {
		fmt.Printf("    visited node: %q\n", n1)
		n1.Status = 1
	}

	getIndex := func(slice []*GNode, node *GNode) int {
		for i := range slice {
			if slice[i] == node {
				return i
			}
		}
		return -1
	}

	mergeSlice := func(s1, s2 []*GNode) (out []*GNode) {
		out = make([]*GNode, 0, len(s1)+len(s2))
		out = append(out, s1...)

		for i := range s2 {
			if getIndex(out, s2[i]) == -1 {
				out = append(out, s2[i])
			}
		}
		return
	}

	///
	for len(queue) > 0 {
		unvisited := []*GNode{}
		for i := range queue {
			visit(queue[i])
			unvisited = mergeSlice(unvisited, queue[i].Unvisited())
		}

		queue = unvisited
	}
}

func InstGNodeBFS1() {
	fmt.Println(">>> InstGNodeBFS1:")

	s, a, d, g := NewGNode("S"), NewGNode("A"), NewGNode("D"), NewGNode("G")
	e, b, f, c := NewGNode("E"), NewGNode("B"), NewGNode("F"), NewGNode("C")

	edges := s.BuildPath(a, d, g, e, b, s) + s.BuildPath(c, f, g)
	fmt.Printf("    ~ number of edges: %d\n", edges/2)

	s.BFS()
}

func InstGNodeBFS2() {
	fmt.Println(">>> InstGNodeBFS2:")

	a, b, c, d := NewGNode("A"), NewGNode("B"), NewGNode("C"), NewGNode("D")
	e, f, g := NewGNode("E"), NewGNode("F"), NewGNode("G")

	edges := a.BuildPath(b, c, d) + a.AddTargets(e) + a.BuildPath(f, g)
	fmt.Printf("    ~ number of edges: %d\n", edges/2)

	a.BFS()
}
