package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"
)

// ? THE GRAPH
type WeightedGraph struct {
	Nodes []*Node
	Edges map[string][]*Edge
	mutex sync.RWMutex
}

// ? A GRAPH EDGE, YOU CAN CREATE EDGES
// ? BY DOING graph.Edges["London"]=Edge{<BRAGA NODE>,<COST>}
type Edge struct {
	node *Node
	cost int
}

// ? This represents a node and also has
// ? The dijkstra table in it, the lowest cost and the previous node
type Node struct {
	name         string
	lowestCost   int
	previousNode *Node
}

func NewGraph() *WeightedGraph {
	return &WeightedGraph{
		Edges: make(map[string][]*Edge),
	}
}

// ? Function to initi the list of nodes
func AddNodes(g *WeightedGraph, names ...string) (nodes map[string]*Node) {
	nodes = make(map[string]*Node)
	for _, name := range names {
		n := &Node{name, math.MaxInt, nil}
		g.AddNode(n)
		nodes[name] = n
	}
	return
}

// ? Function to add the node to the list of nodes of the graph
func (g *WeightedGraph) AddNode(node *Node) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.Nodes = append(g.Nodes, node)
}

// ? Function to add a Edge
func (g *WeightedGraph) AddEdge(node1, node2 *Node, cost int) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.Edges[node1.name] = append(g.Edges[node1.name], &Edge{node: node2, cost: cost})
	g.Edges[node2.name] = append(g.Edges[node2.name], &Edge{node: node1, cost: cost})
}

func (n *Node) String() string {
	return n.name
}

func (e *Edge) String() string {
	return e.node.String() + "(" + strconv.Itoa(e.cost) + ")"
}

func (g *WeightedGraph) String() (s string) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	for _, n := range g.Nodes {
		s = s + n.String() + " ->"
		for _, c := range g.Edges[n.name] {
			s = s + " " + c.node.String() + " (" + strconv.Itoa(c.cost) + ")"
		}
		s = s + "\n"
	}
	return
}

func buildGraph() *WeightedGraph {
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

func main() {
	graph := buildGraph()
	fmt.Println(graph)
}
