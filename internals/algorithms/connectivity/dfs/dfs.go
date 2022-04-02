package dfs

import "edaa/internals/interfaces"

var explored map[string]bool

func DFS(graph interfaces.Graph) (chan interfaces.Edge, chan interface{}) {
	outChan := make(chan interfaces.Edge)
	killChan := make(chan interface{})
	go func() {
		explored = make(map[string]bool)
		nodes := graph.Nodes()

		for _, n := range nodes {
			explored[n.Id()] = false
		}

		dfs(outChan, killChan, nodes[0])
		killChan <- ""
	}()
	return outChan, killChan
}

func dfs(outChan chan interfaces.Edge, killChan chan interface{}, n interfaces.Node) {
	explored[n.Id()] = true
	for _, e := range n.OutEdges() {
		if !explored[e.To().Id()] {
			outChan <- e
			dfs(outChan, killChan, e.To())
		}
	}
}
