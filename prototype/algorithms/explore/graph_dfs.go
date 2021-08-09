package explore

import (
	"fmt"
)

// Dеpth First Trаvеrsаl
func (node *GNode) DFS() {
	var (
		curr  *GNode
		queue []*GNode
	)

	queue = make([]*GNode, 0)

	// push one node to queue, set node to visited and curr to the node
	push := func(n1 *GNode) {
		fmt.Printf("    push node: %q\n", n1)
		n1.Status = 1
		queue = append(queue, n1)
		curr = n1
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
	for push(node); curr != nil; {
		forward := false
		for _, v := range curr.Targets {
			if v.Status == 0 {
				push(v) // curr = queue[len(queue) - 1]
				forward = true
				break
			}
		}

		if !forward {
			pop() // curr = queue[len(queue) - 1]
		}
	}
}

func InstGNodeDFS1() {
	fmt.Println(">>> InstGNodeDFS1:")

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

	s.DFS()
}

func InstGNodeDFS2() {
	fmt.Println(">>> InstGNodeDFS2:")

	a, b, c, d := NewGNode("A"), NewGNode("B"), NewGNode("C"), NewGNode("D")
	e, f, g := NewGNode("E"), NewGNode("F"), NewGNode("G")

	edges := a.BuildPath(b, c, d) + a.AddTargets(e, f) + f.AddTargets(g)
	fmt.Printf("    ~ number of edges: %d\n", edges/2)

	a.DFS()
}
