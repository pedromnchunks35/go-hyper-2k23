package main

import (
	"fmt"
	"sync"
)

type Heap struct {
	elements []*Node
	mutex    sync.RWMutex
}

type Node struct {
	name    string
	value   int
	through *Node
}

func (h *Heap) Size() int {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	return len(h.elements)
}

// ? MAKE THE PUSH OF THE HEAP, RE-ARRANGE IT
func (h *Heap) Push(element *Node) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.elements = append(h.elements, element)
	i := len(h.elements) - 1
	fmt.Printf("The i is %v \n", i)
	fmt.Printf("The parent(i) is %v \n", parent(i))
	for ; h.elements[i].value < h.elements[parent(i)].value; i = parent(i) {
		h.swap(i, parent(i))
		fmt.Printf("[inside] The i is %v \n", i)
		fmt.Printf("[inside] The parent(i) is %v \n", parent(i))
	}
}

func (h *Heap) swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

func parent(i int) int {
	return (i - 1) / 2
}

func main() {
	node := Node{name: "teste", value: 2, through: nil}
	heap := &Heap{}
	heap.Push(&node)
	node2 := Node{name: "teste2", value: 4, through: nil}
	heap.Push(&node2)
	node3 := Node{name: "teste3", value: 3, through: nil}
	heap.Push(&node3)
}
