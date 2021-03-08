package explore

type GNode struct {
	Name    string
	Status  int
	Targets []*GNode
}

func NewGNode(name string, targets ...*GNode) (node *GNode) {
	node = &GNode{Name: name, Status: 0}

	if len(targets) > 0 {
		node.Targets = targets
	} else {
		node.Targets = make([]*GNode, 0)
	}

	return
}

func (node *GNode) String() string {
	if node == nil {
		return "<nil>"
	}
	return node.Name
}

func (node *GNode) addTarget(node2 *GNode) (num int) {
	var ok bool

	ok = false
	for _, v := range node.Targets {
		if v == node2 {
			ok = true
			break
		}
	}
	if !ok {
		num++
		node.Targets = append(node.Targets, node2)
	}

	ok = false
	for _, v := range node2.Targets {
		if v == node {
			ok = true
			break
		}
	}
	if !ok {
		num++
		node2.Targets = append(node2.Targets, node)
	}

	return
}

func (node *GNode) AddTargets(nodes ...*GNode) (num int) {
	for i := range nodes {
		num += node.addTarget(nodes[i])
	}

	return
}
