package graph

type Graph struct {
	nodes []*Node
}

func (g *Graph) Init() {
	g.initBus()
}

func (g *Graph) AddNode(node *Node) {
	g.nodes = append(g.nodes, node)
}

func (g Graph) Draw() {}
