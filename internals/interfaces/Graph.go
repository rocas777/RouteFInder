package interfaces

type Graph interface {
	Nodes() []Node
	SetNodes(nodes []Node)

	NodesMap() map[string]Node
	SetNodesMap(nodesMap map[string]Node)

	WalkableNodes() map[string]Node
	SetWalkableNodes(walkableNodes map[string]Node)

	BusableNodes() map[string]Node
	SetBusableNodes(busableNodes map[string]Node)

	MetroableNodes() map[string]Node
	SetMetroableNodes(metroableNodes map[string]Node)

	MaxLat() float64
	SetMaxLat(maxLat float64)

	MinLat() float64
	SetMinLat(minLat float64)

	MaxLon() float64
	SetMaxLon(maxLon float64)

	MinLon() float64
	SetMinLon(minLon float64)

	GetNode(id string) (Node, error)
	AddNode(node Node)
	RemoveNodes(nodes []Node)
}
