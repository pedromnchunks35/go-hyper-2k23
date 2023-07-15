// Golang Code
package main

import "fmt"

// Node structure to store the data and
// the pointers to the left and
// right children
type Node struct {
	key    int
	left   *Node
	right  *Node
	height int
}

// Function to find max of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Calculates the height of the node
func height(N *Node) int {
	if N == nil {
		return 0
	}
	return N.height
}

// Creates a new Node structure
func newNode(key int) *Node {
	node := &Node{key: key}
	node.left = nil
	node.right = nil
	node.height = 1
	return node
}

// Performs a right rotation on the node
func rightRotate(y *Node) *Node {
	x := y.left
	T2 := x.right
	x.right = y
	y.left = T2
	y.height = max(height(y.left), height(y.right)) + 1
	x.height = max(height(x.left), height(x.right)) + 1
	return x
}

// Performs a left rotation on the node
func leftRotate(x *Node) *Node {
	y := x.right
	T2 := y.left
	y.left = x
	x.right = T2
	x.height = max(height(x.left), height(x.right)) + 1
	y.height = max(height(y.left), height(y.right)) + 1
	return y
}

// Calculates the balance factor
// of the node
func getBalanceFactor(N *Node) int {
	return height(N.left) - height(N.right)
}

// Inserts a new node into the
// AVL Tree
func insertNode(node *Node, key int) *Node {
	if node == nil {
		return newNode(key)
	}
	if key < node.key {
		node.left = insertNode(node.left, key)
	} else if key > node.key {
		node.right = insertNode(node.right, key)
	} else {
		return node
	}

	node.height = 1 + max(height(node.left), height(node.right))
	balanceFactor := getBalanceFactor(node)

	if balanceFactor > 1 {
		if key < node.left.key {
			return rightRotate(node)
		} else if key > node.left.key {
			node.left = leftRotate(node.left)
			return rightRotate(node)
		}
	}

	if balanceFactor < -1 {
		if key > node.right.key {
			return leftRotate(node)
		} else if key < node.right.key {
			node.right = rightRotate(node.right)
			return leftRotate(node)
		}
	}

	return node
}

// Fetches the Node with minimum
// value from the AVL tree
func nodeWithMinimumValue(node *Node) *Node {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current
}

// Deletes a node from the AVL Tree
func deleteNode(root *Node, key int) *Node {

	// Searching Node
	if root == nil {
		return root
	}
	if key < root.key {
		root.left = deleteNode(root.left, key)
	} else if key > root.key {
		root.right = deleteNode(root.right, key)
	} else {
		if root.left == nil || root.right == nil {
			temp := root.left
			if temp == nil {
				temp = root.right
			}
			if temp == nil {
				temp = root
				root = nil
			} else {
				*root = *temp
			}
		} else {
			temp := nodeWithMinimumValue(root.right)
			root.key = temp.key
			root.right = deleteNode(root.right, temp.key)
		}
	}
	if root == nil {
		return root
	}
	root.height = 1 + max(height(root.left), height(root.right))
	balanceFactor := getBalanceFactor(root)

	if balanceFactor > 1 {
		if getBalanceFactor(root.left) >= 0 {
			return rightRotate(root)
		} else {
			root.left = leftRotate(root.left)
			return rightRotate(root)
		}
	}
	if balanceFactor < -1 {
		if getBalanceFactor(root.right) <= 0 {
			return leftRotate(root)
		} else {
			root.right = rightRotate(root.right)
			return leftRotate(root)
		}
	}
	return root
}

// Prints the AVL tree
func printTree(root *Node, indent string, last bool) {
	if root != nil {
		fmt.Print(indent)
		if last {
			fmt.Print("R----")
			indent += " "
		} else {
			fmt.Print("L----")
			indent += "| "
		}
		fmt.Println(root.key)
		printTree(root.left, indent, false)
		printTree(root.right, indent, true)
	}
}

func main() {

	// Creating AVL tree and
	// inserting data in it
	root := insertNode(nil, 33)
	root = insertNode(root, 13)
	root = insertNode(root, 53)
	root = insertNode(root, 9)
	root = insertNode(root, 21)
	root = insertNode(root, 61)
	root = insertNode(root, 8)
	root = insertNode(root, 11)

	// Printing AVL Tree
	printTree(root, "", true)

	// Deleting a node from AVL Tree
	root = deleteNode(root, 13)
	fmt.Println("After deleting ")
	printTree(root, "", true)
}
