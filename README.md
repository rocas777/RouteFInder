# FEUP-EDAA

Program to find a Multimodal path on the city of Porto

![image](https://user-images.githubusercontent.com/28635230/187101672-5e259f66-54c5-4502-b1b8-2bbdff820355.png)


## How to run
go run cmd/main/edaa.go

Open index.html and chose option 2 on the program to use the UI


## Implemented ALgorithms
- A* for Shortest Path
- ALT for Shortest path
- Genetic Algorithm for Shortest and Cheapest Path

## List of dependencies
### Go (automatically installed when executing go run):
- github.com/gocarina/gocsv v0.0.0-20220310154401-d4df709ca055
- github.com/onsi/ginkgo v1.16.5
- github.com/onsi/gomega v1.19.0
- github.com/paulmach/osm v0.2.2
- github.com/starwander/GoFibonacciHeap v0.0.0-20190508061137-ba2e4f01000a
- github.com/umahmood/haversine v0.0.0-20151105152445-808ab04add26
### To visualize the graph v1:
- NetworkX (pip install networkx)
- MatPlotLib (pip install matplotlib)
- Numpy (pip install numpy)

## How to see the graph without using the menu
python networkx/view.py

