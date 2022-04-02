package graph

import (
	"fmt"
	"regexp/syntax"
	"time"
)

type Graph struct {
	Nodes          []*Node
	NodesMap       map[string]*Node
	WalkableNodes  map[string]*Node
	BusableNodes   map[string]*Node
	MetroableNodes map[string]*Node
	maxLat         float64
	minLat         float64
	maxLon         float64
	minLon         float64
}

func NewGraph() *Graph {
	return &Graph{NodesMap: make(map[string]*Node)}
}

func (g *Graph) Init() {
	g.NodesMap = make(map[string]*Node)
	start := time.Now()

	g.initBus()
	println("Bus size:", len(g.Nodes))

	elapsed := time.Since(start)
	fmt.Printf("Bus initiation %s\n\n", elapsed)

	start = time.Now()

	g.initMetro()
	println("Metro+Bus size:", len(g.Nodes))

	elapsed = time.Since(start)
	fmt.Printf("Metro initiation %s\n\n", elapsed)

	start = time.Now()

	g.initRoads()
	println("Roads+Metro+Bus size:", len(g.Nodes))

	elapsed = time.Since(start)
	fmt.Printf("Roads initiation %s\n\n", elapsed)
}

func (g *Graph) fastGetNode(code string) (*Node, error) {
	return g.NodesMap[code], nil
}

func (g *Graph) FilterNodes() {
	nodesCopy := g.Nodes
	g.Nodes = make([]*Node, 0)
	g.NodesMap = make(map[string]*Node)
	for _, node := range nodesCopy {
		if node.Referenced {
			g.AddNode(node)
		}
	}
}

func (g *Graph) slowGetNode(code string) (*Node, error) {
	for _, node := range g.Nodes {
		if node.Code == code {
			return node, nil
		}
	}
	return nil, &syntax.Error{}
}

func (g *Graph) GetNode(code string) (*Node, error) {
	return g.fastGetNode(code)
	return g.slowGetNode(code)
}

func (g *Graph) AddNode(node *Node) {
	g.maxLat = 0
	g.maxLon = 0
	g.minLat = 0
	g.minLon = 0
	g.NodesMap[node.Code] = node
	g.Nodes = append(g.Nodes, node)
}

func (g *Graph) GetEdges() []*Edge {
	var outEdges []*Edge
	for _, node := range g.Nodes {
		outEdges = append(outEdges, node.Edges...)
	}
	return outEdges
}

func (g *Graph) GetCoordsBox() (float64, float64, float64, float64) {
	if g.maxLat == g.maxLon && g.minLat == g.minLon && g.minLat == 0 {
		g.maxLat = -10000
		g.maxLon = -10000
		g.minLat = 10000
		g.minLon = 10000

		for _, node := range g.Nodes {
			if node.latitude >= g.maxLat {
				g.maxLat = node.latitude
			} else if node.latitude <= g.minLat {
				g.minLat = node.latitude
			}

			if node.longitude >= g.maxLon {
				g.maxLon = node.longitude
			} else if node.longitude <= g.minLon {
				g.minLon = node.longitude
			}
		}
	}
	return g.maxLat, g.minLat, g.maxLon, g.minLon
}

// removes every node that does not belong to the biggest SCC and does not include a station
func (g *Graph) RemoveUnconnectedNodes(disconnectedComponents [][]*SCC) {
	biggestScc := 0
	for _, components := range disconnectedComponents {
		for _, component := range components {
			if biggestScc <= len(component.Nodes) {
				biggestScc = len(component.Nodes)
			}
		}
	}
	nodesToBeRemoved := make([]*Node, 0)
	for _, components := range disconnectedComponents {
		for _, component := range components {
			if len(component.Nodes) < biggestScc && !component.HasStation || len(component.Nodes) <= 1 {
				nodesToBeRemoved = append(nodesToBeRemoved, component.Nodes...)
			}
		}
	}
	g.RemoveNodes(nodesToBeRemoved)
}

func (g *Graph) RemoveNodes(nodes []*Node) {
	for _, n := range nodes {
		delete(g.NodesMap, n.Code)
	}
	g.Nodes = make([]*Node, len(g.NodesMap))
	counter := 0
	for _, n := range g.NodesMap {
		g.Nodes[counter] = n
		counter++
	}
}

func (g *Graph) RemoveNode(node *Node) {
	for i, n := range g.Nodes {
		if n.Code == node.Code {
			g.Nodes = append(g.Nodes[:i], g.Nodes[i+1:]...)
			break
		}
	}
	delete(g.NodesMap, node.Code)
}
