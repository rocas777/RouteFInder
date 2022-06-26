package main

import (
	"edaa/internals/algorithms/path/landmarks"
	kdtree2 "edaa/internals/dataStructures/kdtree"
	quadtree2 "edaa/internals/dataStructures/quadtree"
	"edaa/internals/exports/reuse"
	"edaa/internals/graph"
	tile_server "edaa/internals/visualization/tile-server"
	"errors"
	"fmt"
	"log"
	"os/exec"

	//"edaa/internals/algorithms/path/astar"
	//"edaa/internals/g"
	"edaa/internals/interfaces"
	"edaa/internals/menuHelper"

	//"edaa/internals/utils"
	//"time"
	"bufio"
	"os"
	"strings"
)

var g interfaces.Graph
var kdtree *kdtree2.KDTree
var quadtree *quadtree2.QuadTree
var l *landmarks.Dijkstra

var first = 0

var kill chan interface{}


func main() {
	kill = make(chan interface{})
	for {
		fmt.Println("")
		fmt.Println("Select one of the following")
		fmt.Println("1 - Setup")
		fmt.Println("2 - Load")
		fmt.Println("3 - Load Landmarks")
		fmt.Println("4 - Connectivity Analysis")
		fmt.Println("5 - Find Path")
		fmt.Println("6 - Find Path Landmarks")
		fmt.Println("7 - Find Path Genetics")
		fmt.Println("8 - See Map")
		fmt.Println("9 - Load Raw")
		fmt.Println("10 - Export")
		fmt.Println("11 - Exit")
		reader := bufio.NewReader(os.Stdin)
		opt, _ := reader.ReadString('\n')
		opt = strings.TrimSpace(opt)


		switch opt {
		case "1":
			g = &graph.Graph{}
			g = menuHelper.Setup()
		case "2":
			if _, err := os.Stat("data/reuse/edges.csv"); err == nil {
				g = &graph.Graph{}
				graph.InitReuse(g)
			} else if errors.Is(err, os.ErrNotExist) {
				fmt.Println("Could not load, file does not exist")
			}
			if quadtree == nil {
				quadtree = quadtree2.NewQuadTree(g)
				go restartServer()
			}
		case "3":
			if g == nil {
				fmt.Println("Must setup or load graph first!!!!")
			} else {
				l = menuHelper.Landmarks(g)
			}
		case "4":
			if g == nil {
				fmt.Println("Must setup or load graph first!!!!")
			} else {
				menuHelper.Connectivity(g)
			}
		case "5":
			if g == nil {
				fmt.Println("Must setup or load graph first!!!!")
			} else {
				if g == nil {
					fmt.Println("Must setup or load graph first!!!!")
				} else {
					if kdtree == nil {
						kdtree = kdtree2.NewKDTree(g)
					}
					if quadtree == nil {
						quadtree = quadtree2.NewQuadTree(g)
						go restartServer()
					}
					menuHelper.PathFinder(g, kdtree)
				}
			}
		case "6":
			if g == nil {
				fmt.Println("Must setup or load graph first!!!!")
			} else {
				if l == nil {
					fmt.Println("Must load landmarks first!!!!")
				} else {
					if kdtree == nil {
						kdtree = kdtree2.NewKDTree(g)
					}
					if quadtree == nil {
						quadtree = quadtree2.NewQuadTree(g)
						go restartServer()
					}
					menuHelper.PathFinderLandmarks(g, kdtree, l)
				}
			}
		case "7":
			if g == nil {
				fmt.Println("Must setup or load graph first!!!!")
			} else {
				if g == nil {
					fmt.Println("Must setup or load graph first!!!!")
				} else {
					if kdtree == nil {
						kdtree = kdtree2.NewKDTree(g)
					}
					if quadtree == nil {
						quadtree = quadtree2.NewQuadTree(g)
						go restartServer()
					}
					menuHelper.PathFinderGenetics(g, kdtree)
				}
			}
		case "8":
			go func() {
				fmt.Println("Loading... be patient")
				cmd := exec.Command("python3", "networkx/view.py")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				log.Println(cmd.Run())
			}()
		case "9":
			g = &graph.Graph{}
			fmt.Println("")
			fmt.Println("Initiating graph...")
			g.Init()
		case "10":
			reuse.ExportEdges(g, "data/reuse/edges.csv")
			reuse.ExportNodes(g, "data/reuse/nodes.csv")
		case "11":
			fmt.Println("Cya")
			return
		}
	}
}

func restartServer(){
	if first != 0 {
		kill <- ""
	}
	first = 1
	tile_server.TileServer(quadtree.Root,kill)
}