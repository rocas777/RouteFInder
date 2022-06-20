package graph

import (
	"edaa/internals/interfaces"
	"edaa/internals/utils"
	"fmt"
	"regexp/syntax"
	"time"
)

type Graph struct {
	nodes          []interfaces.Node
	nodesMap       map[string]interfaces.Node
	walkableNodes  map[string]interfaces.Node
	busableNodes   map[string]interfaces.Node
	metroableNodes map[string]interfaces.Node
	maxLat         float64
	minLat         float64
	maxLon         float64
	minLon         float64
}

func NewGraph() *Graph {
	return &Graph{nodesMap: make(map[string]interfaces.Node)}
}

func (g *Graph) Nodes() []interfaces.Node {
	return g.nodes
}

func (g *Graph) SetNodes(nodes []interfaces.Node) {
	g.nodes = nodes
}

func (g *Graph) NodesMap() map[string]interfaces.Node {
	return g.nodesMap
}

func (g *Graph) SetNodesMap(nodesMap map[string]interfaces.Node) {
	g.nodesMap = nodesMap
}

func (g *Graph) WalkableNodes() map[string]interfaces.Node {
	return g.walkableNodes
}

func (g *Graph) SetWalkableNodes(walkableNodes map[string]interfaces.Node) {
	g.walkableNodes = walkableNodes
}

func (g *Graph) BusableNodes() map[string]interfaces.Node {
	return g.busableNodes
}

func (g *Graph) SetBusableNodes(busableNodes map[string]interfaces.Node) {
	g.busableNodes = busableNodes
}

func (g *Graph) MetroableNodes() map[string]interfaces.Node {
	return g.metroableNodes
}

func (g *Graph) SetMetroableNodes(metroableNodes map[string]interfaces.Node) {
	g.metroableNodes = metroableNodes
}

func (g *Graph) MaxLat() float64 {
	return g.maxLat
}

func (g *Graph) SetMaxLat(maxLat float64) {
	g.maxLat = maxLat
}

func (g *Graph) MinLat() float64 {
	return g.minLat
}

func (g *Graph) SetMinLat(minLat float64) {
	g.minLat = minLat
}

func (g *Graph) MaxLon() float64 {
	return g.maxLon
}

func (g *Graph) SetMaxLon(maxLon float64) {
	g.maxLon = maxLon
}

func (g *Graph) MinLon() float64 {
	return g.minLon
}

func (g *Graph) SetMinLon(minLon float64) {
	g.minLon = minLon
}

func (g *Graph) Init() {
	g.nodesMap = make(map[string]interfaces.Node)
	g.walkableNodes = make(map[string]interfaces.Node)
	g.busableNodes = make(map[string]interfaces.Node)
	g.metroableNodes = make(map[string]interfaces.Node)
	start := time.Now()

	InitBus(g)
	println("Bus size:", len(g.nodes))

	elapsed := time.Since(start)
	fmt.Printf("Bus initiation %s\n\n", elapsed)

	start = time.Now()

	InitMetro(g)
	println("Metro+Bus size:", len(g.nodes))

	elapsed = time.Since(start)
	fmt.Printf("Metro initiation %s\n\n", elapsed)

	start = time.Now()

	InitRoads(g)
	println("Roads+Metro+Bus size:", len(g.nodes))

	elapsed = time.Since(start)
	fmt.Printf("Roads initiation %s\n\n", elapsed)
}

func (g *Graph) fastGetNode(code string) (interfaces.Node, error) {
	return g.nodesMap[code], nil
}

func (g *Graph) slowGetNode(code string) (interfaces.Node, error) {
	for _, node := range g.nodes {
		if node.Id() == code {
			return node, nil
		}
	}
	return nil, &syntax.Error{}
}

func (g *Graph) GetNode(code string) (interfaces.Node, error) {
	return g.fastGetNode(code)
	return g.slowGetNode(code)
}

func (g *Graph) AddNode(node interfaces.Node) {
	g.maxLat = 0
	g.maxLon = 0
	g.minLat = 0
	g.minLon = 0
	g.nodesMap[node.Id()] = node
	g.nodes = append(g.nodes, node)
}

func (g *Graph) GetEdges() []interfaces.Edge {
	var outEdges []interfaces.Edge
	for _, node := range g.nodes {
		outEdges = append(outEdges, node.OutEdges()...)
	}
	return outEdges
}

func (g *Graph) GetCoordsBox() (float64, float64, float64, float64) {
	if g.maxLat == g.maxLon && g.minLat == g.minLon && g.minLat == 0 {
		g.maxLat = -10000
		g.maxLon = -10000
		g.minLat = 10000
		g.minLon = 10000

		for _, node := range g.nodes {
			if node.Latitude() >= g.maxLat {
				g.maxLat = node.Latitude()
			} else if node.Latitude() <= g.minLat {
				g.minLat = node.Latitude()
			}

			if node.Longitude() >= g.maxLon {
				g.maxLon = node.Longitude()
			} else if node.Longitude() <= g.minLon {
				g.minLon = node.Longitude()
			}
		}
	}
	return g.maxLat, g.minLat, g.maxLon, g.minLon
}

func (g *Graph) RemoveNodes(nodes []interfaces.Node) {
	for _, n := range nodes {
		delete(g.nodesMap, n.Id())
		delete(g.busableNodes, n.Id())
		delete(g.metroableNodes, n.Id())
		delete(g.walkableNodes, n.Id())
	}
	g.nodes = make([]interfaces.Node, len(g.nodesMap))
	counter := 0
	for _, n := range g.nodesMap {
		g.nodes[counter] = n
		counter++
	}
}

func (g *Graph) RemoveNode(node interfaces.Node) {
	for i, n := range g.nodes {
		if n.Id() == node.Id() {
			g.nodes = append(g.nodes[:i], g.nodes[i+1:]...)
			break
		}
	}
	delete(g.nodesMap, node.Id())
}

func (g *Graph) GetClosestNode(node interfaces.Node) (interfaces.Node, float64) {
	closestDistance := 1000000000000000000000.0
	var closestNode interfaces.Node
	for _, wn := range g.walkableNodes {
		dist := utils.GetDistance(node.Latitude(), node.Longitude(), wn.Latitude(), wn.Longitude())
		if dist < closestDistance {
			closestDistance = dist
			closestNode = wn
		}
	}
	return closestNode, closestDistance
}
