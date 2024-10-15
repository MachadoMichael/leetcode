package main

import "fmt"

type Graph struct {
	NumberOfNodes int
	AdjacencyList map[int][]int
}

func (g *Graph) addVertex(node int) {
	g.AdjacencyList[node] = []int{}
	g.NumberOfNodes++
}

func (g *Graph) addEdge(node1 int, node2 int) {
	g.AdjacencyList[node1] = append(g.AdjacencyList[node1], node2)
	g.AdjacencyList[node2] = append(g.AdjacencyList[node2], node1)
}

func (g *Graph) showConnections() {
	println("NumberOfNodes: ", g.NumberOfNodes)
	for k, values := range g.AdjacencyList {
		formatedValues := fmt.Sprintf("%v -> ", k)
		for _, v := range values {
			formatedValues += fmt.Sprintf("%v ", v)
		}
		println(formatedValues)
	}
}

func main() {
	graph := Graph{
		NumberOfNodes: 0,
		AdjacencyList: map[int][]int{},
	}

	graph.addVertex(0)
	graph.addVertex(1)
	graph.addVertex(2)
	graph.addVertex(3)
	graph.addVertex(4)
	graph.addVertex(5)
	graph.addVertex(6)
	graph.addEdge(3, 1)
	graph.addEdge(3, 4)
	graph.addEdge(4, 2)
	graph.addEdge(4, 5)
	graph.addEdge(1, 2)
	graph.addEdge(1, 0)
	graph.addEdge(0, 2)
	graph.addEdge(6, 5)
	graph.showConnections()
}
