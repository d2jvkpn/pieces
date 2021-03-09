package explore

import (
	"fmt"
)

// Brеаdth First Trаvеrsаl
func (node *GNode) BFS() {
	var (
		numOfUnvisited int
		curr           *GNode
		queue          []*GNode
	)

	queue = make([]*GNode, 0)

	// push one node to queue, set node to visited and curr to the node
	push := func(n1 *GNode) {
		fmt.Printf("    push node: %q\n", n1)
		n1.Status = 1
		queue = append(queue, n1)
		// curr = n1
	}

	// pop one node to queue, set curr to the latest one of queue
	pop := func() (n2 *GNode) {
		if len(queue) == 0 {
			curr = nil
			return nil
		}

		n2 = queue[len(queue)-1]
		fmt.Printf("    pop node: %q\n", n2)
		queue = queue[:(len(queue) - 1)]

		if len(queue) > 0 {
			curr = queue[len(queue)-1]
		} else {
			curr = nil
		}
		return
	}

	///
	push(node)
	curr, numOfUnvisited = node, len(node.Unvisited())

	for idx := 0; curr != nil; {
		for _, v := range curr.Unvisited() {
			numOfUnvisited += len(v.Unvisited())
			push(v) // not set curr
			numOfUnvisited--
		}

		if numOfUnvisited > 0 {
			idx++
			curr = queue[idx]
		} else {
			pop() // curr = queue[len(queue) - 1]
		}
	}
}

func InstGNodeBFS1() {
	fmt.Println(">>> InstGNodeBFS1:")

	s, a, d := NewGNode("S"), NewGNode("A"), NewGNode("D")
	g, e, b := NewGNode("G"), NewGNode("E"), NewGNode("B")
	f, c := NewGNode("F"), NewGNode("C")

	edges := 0
	edges += s.AddTargets(a, b, c)
	edges += a.AddTargets(d)
	edges += d.AddTargets(g)
	edges += g.AddTargets(e, f)
	edges += e.AddTargets(b)
	edges += f.AddTargets(c)

	fmt.Printf("    ~ number of edges: %d\n", edges/2)

	s.BFS()
}

func InstGNodeBFS2() {
	fmt.Println(">>> InstGNodeBFS2:")

	a, b, c := NewGNode("A"), NewGNode("B"), NewGNode("C")
	d, e, f := NewGNode("D"), NewGNode("E"), NewGNode("F")
	g := NewGNode("G")

	edges := 0
	edges += a.BuildPath(b, c, d)
	edges += a.AddTargets(e)
	edges += a.BuildPath(f, g)

	fmt.Printf("    ~ number of edges: %d\n", edges/2)

	a.BFS()
}
