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
	"net/http"
	"os/exec"
	"strconv"

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

func server (){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		sLatS := r.URL.Query()["slat"][0]
		sLonS := r.URL.Query()["slon"][0]
		dLatS := r.URL.Query()["dlat"][0]
		dLonS := r.URL.Query()["dlon"][0]
		m := r.URL.Query()["method"][0]

		sLat,_ := strconv.ParseFloat(sLatS,64)
		sLon,_ := strconv.ParseFloat(sLonS,64)
		dLat,_ := strconv.ParseFloat(dLatS,64)
		dLon,_ := strconv.ParseFloat(dLonS,64)

		println(m)
		switch m{
		case "d":
			time,cost := menuHelper.DijkstraServer(g,kdtree,sLat,sLon,dLat,dLon)
			println(time,cost)
		case "a":
			time,cost := menuHelper.AStartServer(g,kdtree,sLat,sLon,dLat,dLon)
			println(time,cost)
		case "alt":
			time,cost := menuHelper.ALTServer(g,kdtree,sLat,sLon,dLat,dLon,l)
			println(time,cost)
		case "gt":
			time,cost := menuHelper.GeneticTimeServer(g,kdtree,sLat,sLon,dLat,dLon)
			println(time,cost)
		case "gp":
			time,cost := menuHelper.GeneticPriceServer(g,kdtree,sLat,sLon,dLat,dLon)
			println(time,cost)
		}

		//fmt.Println(sLat,sLon,dLat,dLon)
		w.Write([]byte(""))
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	go server()
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