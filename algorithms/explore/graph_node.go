package explore

type GNode struct {
	Name    string
	Status  int
	Targets []*GNode
}

func (node *GNode) String() string {
	if node == nil {
		return "<nil>"
	}
	return node.Name
}
