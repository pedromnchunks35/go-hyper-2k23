package main

import "fmt"

// ? CREATION OF GRAPH
type Graph struct {
	vertices int
	adjList  map[int][]int
}

func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: 5,
		adjList:  make(map[int][]int),
	}
}

func (g *Graph) AddEdge(source, dest int) {
	g.adjList[source] = append(g.adjList[source], dest)
	g.adjList[dest] = append(g.adjList[dest], source)
}

// ? SEARCH OVER THE GRAPH USING DFS
func (g *Graph) DFS(startVertex int) {
	visited := make(map[int]bool)
	g.DFSUtil(startVertex, visited)
}

func (g *Graph) DFSUtil(vertex int, visited map[int]bool) {
	visited[vertex] = true
	fmt.Printf("%d \n", vertex)
	for _, v := range g.adjList[vertex] {
		if !visited[v] {
			g.DFSUtil(v, visited)
		}
	}
}

func main() {
	g := NewGraph(5)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.DFS(0)
}
