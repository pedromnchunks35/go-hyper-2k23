package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"sync"
)

type Graph struct {
	Nodes []*Node
	Edges map[string][]*Edge
	mutex sync.RWMutex
}

type Edge struct {
	Node *Node
	Cost int
}

type Node struct {
	Name         string
	MinCost      int
	PreviousNode *Node
}

func NewGraph() *Graph {
	return &Graph{
		Edges: make(map[string][]*Edge),
	}
}

func (g *Graph) GetNode(name string) (node *Node) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	for _, n := range g.Nodes {
		if n.Name == name {
			node = n
			break
		}
	}
	return
}

func (g *Graph) AddNode(n *Node) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.Nodes = append(g.Nodes, n)
}

func AddNodes(g *Graph, names ...string) (nodes map[string]*Node) {
	nodes = make(map[string]*Node)
	for _, name := range names {
		n := &Node{Name: name, MinCost: math.MaxInt, PreviousNode: nil}
		g.AddNode(n)
		nodes[name] = n
	}
	return
}

func (g Graph) AddEdge(node1, node2 *Node, cost int) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.Edges[node1.Name] = append(g.Edges[node1.Name], &Edge{Node: node2, Cost: cost})
	g.Edges[node2.Name] = append(g.Edges[node2.Name], &Edge{Node: node1, Cost: cost})
}

func (n *Node) String() string {
	return n.Name
}

func (e *Edge) String() string {
	return e.Node.String() + "[" + strconv.Itoa(e.Cost) + "]"
}

func (g *Graph) String() (s string) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	for _, n := range g.Nodes {
		s += n.String() + "->"
		for _, e := range g.Edges[n.Name] {
			s += " " + e.String()
		}
		s += "\n"
	}
	return
}

func buildGraph() *Graph {
	graph := NewGraph()
	nodes := AddNodes(graph,
		"Braga",
		"Arcos",
		"Porto",
		"Leiria",
		"Lisboa",
	)
	graph.AddEdge(nodes["Braga"], nodes["Porto"], 5)
	graph.AddEdge(nodes["Braga"], nodes["Arcos"], 2)
	graph.AddEdge(nodes["Arcos"], nodes["Porto"], 6)
	graph.AddEdge(nodes["Arcos"], nodes["Leiria"], 2)
	graph.AddEdge(nodes["Porto"], nodes["Leiria"], 2)
	graph.AddEdge(nodes["Porto"], nodes["Lisboa"], 3)
	graph.AddEdge(nodes["Leiria"], nodes["Lisboa"], 4)
	return graph
}

func getEdgeWithMinimumCost(edges []*Edge, isVisited *map[string]bool) (minEdge *Edge) {
	for _, edge := range edges {
		if (minEdge == nil || edge.Cost < minEdge.Cost) && (*isVisited)[edge.Node.Name] != true {
			minEdge = edge
		}
	}
	return
}

func (g *Graph) dijstra(name string) {
	//? First member
	isVisited := make(map[string]bool)
	current := g.GetNode(name)
	current.MinCost = 0
	isVisited[name] = true
	edges := g.Edges[name]
	//? The first edges will be putted on the map without objections
	for i := 0; i < len(edges); i++ {
		edges[i].Node.MinCost = edges[i].Cost
		edges[i].Node.PreviousNode = current
	}
	//? We will gather the next edge with the minimum cost
	nextEdge := getEdgeWithMinimumCost(edges, &isVisited)
	//? Loop, passing by every nodes until they are all visited and the costs
	//? from a certain point reduce
	var oldCost int
	var newCost int
	for len(isVisited) != len(g.Nodes) {
		current = nextEdge.Node
		isVisited[current.Name] = true
		edges = g.Edges[current.Name]
		//? Sum the cost associated with a chain (The minCost until our current node + the cost of the node we are acessing)
		//? Check if it is cheaper,
		//? case it is, change the minCost and the previous node of that node
		for _, edge := range edges {
			oldCost = edge.Node.MinCost
			newCost = current.MinCost + edge.Cost
			if newCost < oldCost {
				edge.Node.MinCost = newCost
				edge.Node.PreviousNode = current
			}
			nextEdge = getEdgeWithMinimumCost(edges, &isVisited)
		}
	}
}

func main() {
	city := os.Args[1]
	graph := buildGraph()
	graph.dijstra(city)
	// display the nodes
	for _, node := range graph.Nodes {
		fmt.Printf("Shortest time from %s to %s is %d\n",
			city, node.Name, node.MinCost)
		for n := node; n.PreviousNode != nil; n = n.PreviousNode {
			fmt.Print(n, " <- ")
		}
		fmt.Println(city)
		fmt.Println()
	}
}
