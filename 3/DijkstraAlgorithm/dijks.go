package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"sync"
)

func main() {
	// build and run Dijkstra's algorithm on graph
	graph := buildGraph()
	city := os.Args[1]
	dijkstra(graph, city)

	// display the nodes
	for _, node := range graph.Nodes {
		fmt.Printf("Shortest time from %s to %s is %d\n",
			city, node.name, node.value)
		for n := node; n.through != nil; n = n.through {
			fmt.Print(n, " <- ")
		}
		fmt.Println(city)
		fmt.Println()
	}
}

func dijkstra(graph *WeightedGraph, city string) {
	visited := make(map[string]bool)
	heap := &Heap{}

	startNode := graph.GetNode(city)
	startNode.value = 0
	heap.Push(startNode)

	for heap.Size() > 0 {
		current := heap.Pop()
		visited[current.name] = true
		edges := graph.Edges[current.name]
		for _, edge := range edges {
			if !visited[edge.node.name] {
				heap.Push(edge.node)
				if current.value+edge.weight < edge.node.value {
					edge.node.value = current.value + edge.weight
					edge.node.through = current
				}
			}
		}
	}
}

func buildGraph() *WeightedGraph {
	graph := NewGraph()
	nodes := AddNodes(graph,
		"London",
		"Paris",
		"Amsterdam",
		"Luxembourg",
		"Zurich",
		"Rome",
		"Berlin",
		"Vienna",
		"Warsaw",
		"Istanbul",
	)
	graph.AddEdge(nodes["London"], nodes["Paris"], 80)
	graph.AddEdge(nodes["London"], nodes["Luxembourg"], 75)
	graph.AddEdge(nodes["London"], nodes["Amsterdam"], 75)
	graph.AddEdge(nodes["Paris"], nodes["Luxembourg"], 60)
	graph.AddEdge(nodes["Paris"], nodes["Rome"], 125)
	graph.AddEdge(nodes["Luxembourg"], nodes["Berlin"], 90)
	graph.AddEdge(nodes["Luxembourg"], nodes["Zurich"], 60)
	graph.AddEdge(nodes["Luxembourg"], nodes["Amsterdam"], 55)
	graph.AddEdge(nodes["Zurich"], nodes["Vienna"], 80)
	graph.AddEdge(nodes["Zurich"], nodes["Rome"], 90)
	graph.AddEdge(nodes["Zurich"], nodes["Berlin"], 85)
	graph.AddEdge(nodes["Berlin"], nodes["Amsterdam"], 85)
	graph.AddEdge(nodes["Berlin"], nodes["Vienna"], 75)
	graph.AddEdge(nodes["Vienna"], nodes["Rome"], 100)
	graph.AddEdge(nodes["Vienna"], nodes["Istanbul"], 130)
	graph.AddEdge(nodes["Warsaw"], nodes["Berlin"], 80)
	graph.AddEdge(nodes["Warsaw"], nodes["Istanbul"], 180)
	graph.AddEdge(nodes["Rome"], nodes["Istanbul"], 155)

	return graph
}

// -- Weighted Graph

type Node struct {
	name    string
	value   int
	through *Node
}

type Edge struct {
	node   *Node
	weight int
}

type WeightedGraph struct {
	Nodes []*Node
	Edges map[string][]*Edge
	mutex sync.RWMutex
}

func NewGraph() *WeightedGraph {
	return &WeightedGraph{
		Edges: make(map[string][]*Edge),
	}
}

func (g *WeightedGraph) GetNode(name string) (node *Node) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	for _, n := range g.Nodes {
		if n.name == name {
			node = n
		}
	}
	return
}

func (g *WeightedGraph) AddNode(n *Node) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.Nodes = append(g.Nodes, n)
}

func AddNodes(graph *WeightedGraph, names ...string) (nodes map[string]*Node) {
	nodes = make(map[string]*Node)
	for _, name := range names {
		n := &Node{name, math.MaxInt, nil}
		graph.AddNode(n)
		nodes[name] = n
	}
	return
}

func (g *WeightedGraph) AddEdge(n1, n2 *Node, weight int) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.Edges[n1.name] = append(g.Edges[n1.name], &Edge{n2, weight})
	g.Edges[n2.name] = append(g.Edges[n2.name], &Edge{n1, weight})
}

func (n *Node) String() string {
	return n.name
}

func (e *Edge) String() string {
	return e.node.String() + "(" + strconv.Itoa(e.weight) + ")"
}

func (g *WeightedGraph) String() (s string) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	for _, n := range g.Nodes {
		s = s + n.String() + " ->"
		for _, c := range g.Edges[n.name] {
			s = s + " " + c.node.String() + " (" + strconv.Itoa(c.weight) + ")"
		}
		s = s + "\n"
	}
	return
}

// -- Heap

type Heap struct {
	elements []*Node
	mutex    sync.RWMutex
}

func (h *Heap) Size() int {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	return len(h.elements)
}

// push an element to the heap, re-arrange the heap
func (h *Heap) Push(element *Node) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.elements = append(h.elements, element)
	i := len(h.elements) - 1
	for ; h.elements[i].value < h.elements[parent(i)].value; i = parent(i) {
		h.swap(i, parent(i))
	}
}

// pop the top of the heap, which is the min value
func (h *Heap) Pop() (i *Node) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	i = h.elements[0]
	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]
	h.rearrange(0)
	return
}

// rearrange the heap
func (h *Heap) rearrange(i int) {
	smallest := i
	left, right, size := leftChild(i), rightChild(i), len(h.elements)
	if left < size && h.elements[left].value < h.elements[smallest].value {
		smallest = left
	}
	if right < size && h.elements[right].value < h.elements[smallest].value {
		smallest = right
	}
	if smallest != i {
		h.swap(i, smallest)
		h.rearrange(smallest)
	}
}

func (h *Heap) swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

func (h *Heap) String() (str string) {
	return fmt.Sprintf("%q\n", getNames(h.elements))
}

func getNames(nodes []*Node) (names []string) {
	for _, node := range nodes {
		names = append(names, node.name)
	}
	return
}
