package main

import (
	"fmt"
	"strconv"
	"sync"
)

type Graph struct {
	Nodes []*Node
	Edges map[string][]*Edge
	mutex sync.RWMutex
}

type Node struct {
	Name         string
	PreviousNode *Node
}

type Edge struct {
	Node *Node
	Cost int
}

func CreateGraph() (graph *Graph) {
	graph = &Graph{Edges: make(map[string][]*Edge)}
	return
}

func (g *Graph) AddNode(node *Node) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.Nodes = append(g.Nodes, node)
}

func (g *Graph) AddNodes(names ...string) (nodes []*Node) {
	for _, name := range names {
		node := &Node{Name: name, PreviousNode: nil}
		nodes = append(nodes, node)
		g.AddNode(node)
	}
	return
}

func (g *Graph) GetNode(name string) (result *Node) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	for _, node := range g.Nodes {
		if node.Name == name {
			result = node
			break
		}
	}
	return
}

func (g *Graph) addEdge(node1, node2 *Node, cost int) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.Edges[node1.Name] = append(g.Edges[node1.Name], &Edge{Node: node2, Cost: cost})
	g.Edges[node2.Name] = append(g.Edges[node2.Name], &Edge{Node: node1, Cost: cost})
}

func (node Node) String() (s string) {
	s = node.Name
	return
}

func (edge Edge) String() (s string) {
	s += "[" + edge.Node.String() + "]cost: " + strconv.Itoa(edge.Cost)
	return
}

func (g Graph) String() (s string) {
	g.mutex.RLock()
	g.mutex.RUnlock()
	for _, node := range g.Nodes {
		s += node.String() + "<-"
		for _, edge := range g.Edges[node.Name] {
			s += " " + edge.Node.String()
		}
		s += "\n"
	}
	return
}

func buildGraph() (g *Graph) {
	g = CreateGraph()
	g.AddNodes(
		"Viana",
		"Braga",
		"Porto",
		"Arcos",
		"Leiria",
		"Lisboa",
		"Algarve",
	)
	g.addEdge(g.GetNode("Viana"), g.GetNode("Braga"), 1)
	g.addEdge(g.GetNode("Braga"), g.GetNode("Arcos"), 2)
	g.addEdge(g.GetNode("Braga"), g.GetNode("Porto"), 5)
	g.addEdge(g.GetNode("Porto"), g.GetNode("Arcos"), 6)
	g.addEdge(g.GetNode("Porto"), g.GetNode("Leiria"), 2)
	g.addEdge(g.GetNode("Porto"), g.GetNode("Lisboa"), 3)
	g.addEdge(g.GetNode("Leiria"), g.GetNode("Arcos"), 2)
	g.addEdge(g.GetNode("Leiria"), g.GetNode("Lisboa"), 4)
	g.addEdge(g.GetNode("Lisboa"), g.GetNode("Algarve"), 5)
	g.addEdge(g.GetNode("Porto"), g.GetNode("Algarve"), 20)
	return
}

func getTheMinimumEdge(edges []*Edge, isVisited *map[string]bool) (result *Edge) {
	for _, edge := range edges {
		if (result == nil || result.Cost > edge.Cost) && (*isVisited)[edge.Node.Name] != true {
			result = edge
		}
	}
	return
}

func getTheMinimumEdgeEvenVisited(edges []*Edge) (result *Edge) {
	for _, edge := range edges {
		if result == nil || result.Cost > edge.Cost {
			result = edge
		}
	}
	return
}

func getNonVisited(nodes []*Node, isVisited *map[string]bool) (result *Node) {
	for _, node := range nodes {
		if (*isVisited)[node.Name] != true {
			result = node
			break
		}
	}
	return
}

func (g Graph) Prims(name string) (result []*Node) {
	//? ARBITRARY NODE PART
	var nonVisitedNode *Node
	isVisited := make(map[string]bool)
	edges := g.Edges[name]
	current := g.GetNode(name)
	isVisited[name] = true
	nextEdge := getTheMinimumEdge(edges, &isVisited)
	result = append(result, current)
	//? NEXT ONES
	for len(isVisited) != len(g.Nodes) {
		if nextEdge == nil {
			nonVisitedNode = getNonVisited(g.Nodes, &isVisited)
			nextEdge = getTheMinimumEdge(g.Edges[nonVisitedNode.Name], &isVisited)
			if nextEdge == nil {
				nextEdge = getTheMinimumEdgeEvenVisited(g.Edges[nonVisitedNode.Name])
			}
			nextEdge.Node.PreviousNode = nonVisitedNode
			isVisited[nonVisitedNode.Name] = true
			current = nextEdge.Node
		} else {
			nextEdge.Node.PreviousNode = current
			current = nextEdge.Node
		}
		result = append(result, current)
		edges = g.Edges[current.Name]
		isVisited[current.Name] = true
		nextEdge = getTheMinimumEdge(edges, &isVisited)
	}
	return
}

func main() {
	g := buildGraph()
	result := g.Prims("Braga")
	for _, node := range result {
		fmt.Printf(" %v->%v", node.PreviousNode, node)
	}
	fmt.Printf("\n")
}
