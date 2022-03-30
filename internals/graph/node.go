package graph

type Node struct {
	edges     []*Edge
	latitude  float64
	longitude float64
	isStation bool
	name      string
	zone      string
	code      string
}

func NewStationNode(latitude float64, longitude float64, name string, zone string, code string) *Node {
	return &Node{latitude: latitude, longitude: longitude, name: name, zone: zone, code: code, isStation: true}
}

func NewNormalNode(latitude float64, longitude float64, name string, zone string, code string) *Node {
	return &Node{latitude: latitude, longitude: longitude, name: name, zone: zone, code: code, isStation: false}
}

func (n *Node) AddDestination(destination *Node, weight float64) {
	n.edges = append(n.edges, NewEdge(n, destination, weight))
}
