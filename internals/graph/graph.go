package graph

type Graph struct {
	Nodes  []*Node
	maxLat float64
	minLat float64
	maxLon float64
	minLon float64
}

func (g *Graph) Init() {
	g.initBus()
	println(len(g.Nodes))
	g.initMetro()
	println(len(g.Nodes))
}

func (g *Graph) AddNode(node *Node) {
	g.maxLat = 0
	g.maxLon = 0
	g.minLat = 0
	g.minLon = 0
	g.Nodes = append(g.Nodes, node)
}

// returns maxLat, minLat, maxLon, minLon
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

/*func (g *Graph) GetNode(code string) *Node {
	return g.Nodes[code]
}*/

func (g Graph) Draw() {}
