package graph

import (
	"edaa/internals/utils"
	"fmt"
)

type FibHeapNode struct {
	Key				float64
	Code			string
	Root			bool
	Loser			bool
	Parent			string
	Children		[]string
	Degree			int
}

func NewFibHeapNode (key float64, code string) *FibHeapNode {
	return &FibHeapNode{Key: key, Code: code, Root: true, Loser: false, Degree: 0, Parent: ""}
}

func (n *FibHeapNode) BecomeRoot() {
	n.Root = true
	n.Loser = false
	n.Parent = ""
}

func (n *FibHeapNode) BecomeChild(parent string) {
	n.Root = false
	n.Parent = parent
}

func (n *FibHeapNode) AddChild(child string, childDegree int) {
	if (child == n.Code) {
		fmt.Println("tried to child itself", child)
		fmt.Println(n.Children[16])
	}
	n.Children = append(n.Children, child)
	if n.Degree < childDegree+1 {
		n.Degree = childDegree+1
	}
}

func (n *FibHeapNode) RemoveChild(child string) {
	n.Children = utils.RemoveString(n.Children, child);
}

func (n *FibHeapNode) PrintNode(heap *FibHeap) {
	fmt.Println("root:", n.Root, "parent:", n.Parent, "code:", n.Code, "k:", n.Key, "l:", n.Loser, "d:", n.Degree, "children:", n.Children)
	for _, c := range n.Children {
		heap.NodesMap[c].PrintNode(heap)
	}
}