package graph

type Graph struct {
	Nodes []*Node
}

func (g *Graph) Init() {
	g.initBus()
}

func (g *Graph) AddNode(node *Node) {
	g.Nodes = append(g.Nodes, node)
}

/*func (g *Graph) GetNode(code string) *Node {
	return g.Nodes[code]
}*/

func (g Graph) Draw() {}
