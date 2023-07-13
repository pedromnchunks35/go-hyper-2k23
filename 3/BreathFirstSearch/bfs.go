package main

import "fmt"

type Graph struct {
	value    int
	Children []*Graph
}

func (g *Graph) CreateGraph() {
	g10 := Graph{value: 10, Children: nil}
	g9 := Graph{value: 9, Children: nil}
	g5 := Graph{value: 5, Children: make([]*Graph, 2)}
	g5.Children[0] = &g9
	g5.Children[1] = &g10
	g6 := Graph{value: 6, Children: nil}
	g2 := Graph{value: 2, Children: make([]*Graph, 2)}
	g2.Children[0] = &g5
	g2.Children[1] = &g6

	g3 := Graph{value: 3, Children: nil}

	g8 := Graph{value: 8, Children: nil}
	g7 := Graph{value: 7, Children: nil}
	g4 := Graph{value: 4, Children: make([]*Graph, 2)}
	g4.Children[0] = &g7
	g4.Children[1] = &g8

	g.Children[0] = &g2
	g.Children[1] = &g3
	g.Children[2] = &g4
}

func (g *Graph) bfs() []int {
	result := []int{}
	queue := []*Graph{g}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current.value)
		for _, child := range current.Children {
			queue = append(queue, child)
		}
	}
	return result
}

func main() {
	g := Graph{value: 1, Children: make([]*Graph, 3)}
	g.CreateGraph()
	result := g.bfs()
	fmt.Printf("This is the array resulted by the queue: %v\n", result)
}
