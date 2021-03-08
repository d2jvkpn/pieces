package explore

import (
	"fmt"
)

func (node *GNode) Transversal() {
	var (
		curr  *GNode
		queue []*GNode
	)

	fmt.Println(">>> Transverse:")

	queue = make([]*GNode, 0)
	push := func(n1 *GNode) {
		fmt.Printf("    push node: %q\n", curr)
		n1.Status = 1
		queue = append(queue, n1)
		curr = n1
	}

	pop := func() (n2 *GNode) {
		if len(queue) == 0 {
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
	curr = node

LOOP:
	for curr != nil {
		for _, v := range curr.Targets {
			if v.Status == 0 {
				push(v)
				continue LOOP
			}
		}

		if len(queue) == 0 {
			break
		}

		pop()
	}
}
