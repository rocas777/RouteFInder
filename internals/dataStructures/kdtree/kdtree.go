package kdtree

import (
	"edaa/internals/graph"
	"edaa/internals/interfaces"
	"edaa/internals/utils"
	"math"
	"sort"
)

const k = 2

type Node struct {
	LeftNode    *Node
	RightNode   *Node
	CurrentNode *graph.Node
}

func NewNode(leftNode *Node, rightNode *Node, currentNode *graph.Node) *Node {
	return &Node{LeftNode: leftNode, RightNode: rightNode, CurrentNode: currentNode}
}

type KDTree struct {
	Root *Node
}

func NewKDTree(g *graph.Graph) *KDTree {
	kd := &KDTree{}
	walkSlice := make([]*graph.Node, len(g.WalkableNodes))
	i := 0
	for _, n := range g.WalkableNodes {
		walkSlice[i] = n
		i++
	}
	root := <-kd.parallelkdtree(walkSlice, 0)
	//root := kd.slowkdtree(g.Nodes, 0)
	kd.Root = root
	return kd
}

func (kd *KDTree) parallelkdtree(pointList []*graph.Node, depth int) chan *Node {
	outChan := make(chan *Node)
	if depth <= 4 {
		go func() {
			if len(pointList) == 0 {
				outChan <- nil
				return
			}
			axis := depth % k
			sort.Slice(pointList, func(i, j int) bool {
				if axis == 0 {
					return pointList[i].Latitude < pointList[j].Latitude
				}
				return pointList[i].Longitude < pointList[j].Longitude
			})
			median := len(pointList) / 2
			leftChan := kd.parallelkdtree(pointList[:median], depth+1)
			rightChan := kd.parallelkdtree(pointList[median+1:], depth+1)

			current := pointList[median]
			outChan <- NewNode(<-leftChan, <-rightChan, current)
		}()
	} else {
		go func() {
			if len(pointList) == 0 {
				outChan <- nil
				return
			}
			axis := depth % k
			sort.Slice(pointList, func(i, j int) bool {
				if axis == 0 {
					return pointList[i].Latitude < pointList[j].Latitude
				}
				return pointList[i].Longitude < pointList[j].Longitude
			})
			median := len(pointList) / 2
			left := kd.slowkdtree(pointList[:median], depth+1)
			right := kd.slowkdtree(pointList[median+1:], depth+1)

			current := pointList[median]
			outChan <- NewNode(left, right, current)
		}()
	}
	return outChan
}

func (kd *KDTree) slowkdtree(pointList []*graph.Node, depth int) *Node {
	if len(pointList) == 0 {
		return nil
	}
	axis := depth % k
	sort.Slice(pointList, func(i, j int) bool {
		if axis == 0 {
			return pointList[i].Latitude < pointList[j].Latitude
		}
		return pointList[i].Longitude < pointList[j].Longitude
	})
	median := len(pointList) / 2
	left := kd.slowkdtree(pointList[:median], depth+1)
	right := kd.slowkdtree(pointList[median+1:], depth+1)

	current := pointList[median]
	out := NewNode(left, right, current)
	return out
}

type BestEstimation struct {
	x          float64
	y          float64
	estimation float64
	n          *graph.Node
}

func (b *BestEstimation) dist(target interfaces.Node) float64 {
	b.estimation = math.Sqrt((b.x-target.Lat())*(b.x-target.Lat()) + (b.y-target.Lon())*(b.y-target.Lon()))
	return b.estimation
	return utils.GetDistance(
		b.x,
		b.y,
		target.Lat(),
		target.Lon(),
	)
}

func (kd *KDTree) GetClosest(target interfaces.Node) (interfaces.Node, float64) {
	out := kd.parallelSearch(kd.Root, target, 0)
	bestEstimation := out
	return bestEstimation.n, bestEstimation.estimation
}

func (kd *KDTree) parallelSearch(node *Node, target interfaces.Node, depth int) *BestEstimation {
	if node == nil {
		return &BestEstimation{
			x:          10000000,
			y:          10000000,
			estimation: 0,
			n:          nil,
		}
	}
	leftChan := make(chan *BestEstimation)
	rightChan := make(chan *BestEstimation)
	go func() {
		if depth <= 3 {
			leftChan <- kd.parallelSearch(node.LeftNode, target, depth+1)
		} else {
			leftChan <- kd.search(node.LeftNode, target)
		}
	}()

	go func() {
		if depth <= 3 {
			rightChan <- kd.parallelSearch(node.RightNode, target, depth+1)
		} else {
			rightChan <- kd.search(node.RightNode, target)
		}
	}()
	currentEstimation := &BestEstimation{
		x:          node.CurrentNode.Latitude,
		y:          node.CurrentNode.Longitude,
		estimation: 0,
		n:          node.CurrentNode,
	}
	left := <-leftChan
	right := <-rightChan

	if left.dist(target) < currentEstimation.dist(target) {
		currentEstimation = left
	}
	if right.dist(target) < currentEstimation.dist(target) {
		currentEstimation = right
	}

	return currentEstimation
}

func (kd *KDTree) search(node *Node, target interfaces.Node) *BestEstimation {
	if node == nil {
		return &BestEstimation{
			x:          10000000,
			y:          10000000,
			estimation: 0,
			n:          nil,
		}
	}
	left := kd.search(node.LeftNode, target)
	right := kd.search(node.RightNode, target)
	currentEstimation := &BestEstimation{
		x:          node.CurrentNode.Latitude,
		y:          node.CurrentNode.Longitude,
		estimation: 0,
		n:          node.CurrentNode,
	}

	if left.dist(target) < currentEstimation.dist(target) {
		currentEstimation = left
	}
	if right.dist(target) < currentEstimation.dist(target) {
		currentEstimation = right
	}

	return currentEstimation
}

func (kd *KDTree) getBestEstimation(estimation *BestEstimation, target *Node, currentDist float64, currentEstimation *BestEstimation) (float64, *BestEstimation) {
	dist := utils.GetDistance(
		estimation.x,
		estimation.y,
		target.CurrentNode.Latitude,
		target.CurrentNode.Longitude,
	)
	if dist < currentDist {
		currentEstimation = estimation
	}
	return dist, currentEstimation
}
