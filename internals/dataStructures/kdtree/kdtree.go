package kdtree

import (
	"edaa/internals/interfaces"
	"edaa/internals/utils"
	"math"
	"sort"
)

const k = 2

type KDTree struct {
	root *Node
}

type Node struct {
	leftNode    *Node
	rightNode   *Node
	currentNode interfaces.Node
}

type BestEstimation struct {
	x          float64
	y          float64
	estimation float64
	n          interfaces.Node
}

func NewNode(leftNode *Node, rightNode *Node, currentNode interfaces.Node) *Node {
	return &Node{leftNode: leftNode, rightNode: rightNode, currentNode: currentNode}
}

func NewKDTree(g interfaces.Graph) *KDTree {
	kd := &KDTree{}
	walkSlice := make([]interfaces.Node, len(g.WalkableNodes()))
	i := 0
	for _, n := range g.WalkableNodes() {
		walkSlice[i] = n
		i++
	}
	root := <-kd.parallelkdtree(walkSlice, 0)
	//root := kd.slowkdtree(g.nodes, 0)
	kd.root = root
	return kd
}

func (kd *KDTree) parallelkdtree(pointList []interfaces.Node, depth int) chan *Node {
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
					return pointList[i].Latitude() < pointList[j].Latitude()
				}
				return pointList[i].Longitude() < pointList[j].Longitude()
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
					return pointList[i].Latitude() < pointList[j].Latitude()
				}
				return pointList[i].Longitude() < pointList[j].Longitude()
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

func (kd *KDTree) slowkdtree(pointList []interfaces.Node, depth int) *Node {
	if len(pointList) == 0 {
		return nil
	}
	axis := depth % k
	sort.Slice(pointList, func(i, j int) bool {
		if axis == 0 {
			return pointList[i].Latitude() < pointList[j].Latitude()
		}
		return pointList[i].Longitude() < pointList[j].Longitude()
	})
	median := len(pointList) / 2
	left := kd.slowkdtree(pointList[:median], depth+1)
	right := kd.slowkdtree(pointList[median+1:], depth+1)

	current := pointList[median]
	out := NewNode(left, right, current)
	return out
}

func (b *BestEstimation) dist(target interfaces.Node) float64 {
	b.estimation = math.Sqrt((b.x-target.Latitude())*(b.x-target.Latitude()) + (b.y-target.Longitude())*(b.y-target.Longitude()))
	return b.estimation
	return utils.GetDistance(
		b.x,
		b.y,
		target.Latitude(),
		target.Longitude(),
	)
}

func (kd *KDTree) GetClosest(target interfaces.Node) (interfaces.Node, float64) {
	out := kd.parallelSearch(kd.root, target, 0)
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
			leftChan <- kd.parallelSearch(node.leftNode, target, depth+1)
		} else {
			leftChan <- kd.search(node.leftNode, target, depth+1)
		}
	}()

	go func() {
		if depth <= 3 {
			rightChan <- kd.parallelSearch(node.rightNode, target, depth+1)
		} else {
			rightChan <- kd.search(node.rightNode, target, depth+1)
		}
	}()
	currentEstimation := &BestEstimation{
		x:          node.currentNode.Latitude(),
		y:          node.currentNode.Longitude(),
		estimation: 0,
		n:          node.currentNode,
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

func (kd *KDTree) search(node *Node, target interfaces.Node, depth int) *BestEstimation {
	if node == nil {
		return &BestEstimation{
			x:          10000000,
			y:          10000000,
			estimation: 0,
			n:          nil,
		}
	}
	currentEstimation := &BestEstimation{
		x:          node.currentNode.Latitude(),
		y:          node.currentNode.Longitude(),
		estimation: 0,
		n:          node.currentNode,
	}
	currentEstimation.dist(target)

	axis := depth % k
	if axis == 0 {
		left := kd.search(node.leftNode, target, depth+1)
		left.dist(target)

		if left.estimation < currentEstimation.estimation {
			currentEstimation = left
		}
		if math.Abs(currentEstimation.x-left.x) < currentEstimation.estimation {
			right := kd.search(node.rightNode, target, depth+1)
			right.dist(target)

			if right.estimation < currentEstimation.estimation {
				currentEstimation = right
			}
		}
	} else {
		right := kd.search(node.rightNode, target, depth+1)
		right.dist(target)

		if right.estimation < currentEstimation.estimation {
			currentEstimation = right
		}
		if math.Abs(currentEstimation.y-right.y) < currentEstimation.estimation {
			left := kd.search(node.leftNode, target, depth+1)
			left.dist(target)

			if left.estimation < currentEstimation.estimation {
				currentEstimation = left
			}
		}
	}

	return currentEstimation
}

func (kd *KDTree) getBestEstimation(estimation *BestEstimation, target *Node, currentDist float64, currentEstimation *BestEstimation) (float64, *BestEstimation) {
	dist := utils.GetDistance(
		estimation.x,
		estimation.y,
		target.currentNode.Latitude(),
		target.currentNode.Longitude(),
	)
	if dist < currentDist {
		currentEstimation = estimation
	}
	return dist, currentEstimation
}
