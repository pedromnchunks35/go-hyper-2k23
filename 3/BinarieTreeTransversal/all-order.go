package main

import (
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (node *Node) Insert(value int) {
	for node != nil {

		if node.Value > value {
			if node.Left == nil {
				node.Left = &Node{Value: value}
				break
			}
			node = node.Left
		} else {
			if node.Right == nil {
				node.Right = &Node{Value: value}
				break
			}
			node = node.Right
		}
	}
}

func preOrderTraversal(node *Node) {
	if node != nil {
		fmt.Printf("%d ", node.Value)
		preOrderTraversal(node.Left)
		preOrderTraversal(node.Right)
	}
}

func inOrderTraversal(node *Node) {
	if node != nil {
		inOrderTraversal(node.Left)
		fmt.Printf("%d ", node.Value)
		inOrderTraversal(node.Right)
	}
}

func postOrderTraversal(node *Node) {
	if node != nil {
		postOrderTraversal(node.Left)
		postOrderTraversal(node.Right)
		fmt.Printf("%d ", node.Value)
	}
}

func main() {
	// Create the binary tree
	root := &Node{
		Value: 5,
		Left: &Node{
			Value: 4,
			Left: &Node{
				Value: 2,
			},
			Right: &Node{
				Value: 1,
			},
		},
		Right: &Node{
			Value: 3,
		},
	}
	// Traverse the tree using pre-order traversal
	fmt.Println("Pre-order traversal:")
	preOrderTraversal(root)

	// Traverse the tree using in-order traversal
	fmt.Println("\nIn-order traversal:")
	inOrderTraversal(root)

	// Traverse the tree using post-order traversal
	fmt.Println("\nPost-order traversal:")
	postOrderTraversal(root)
}
