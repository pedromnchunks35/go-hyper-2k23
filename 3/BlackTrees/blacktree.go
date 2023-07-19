package main

import "fmt"

type Node struct {
	Value       int
	Color       int
	Left, Right *Node
	Parent      *Node
}

func NewNode(value int) (result *Node) {
	return &Node{Value: value, Color: 0, Left: nil, Right: nil, Parent: nil}
}

func (node *Node) GetValue() (result int) {
	result = node.Value
	return
}

/*
* @receiver node, the node we are analysing
* @returns black case the color = 1, red case the color=0
 */
func (node *Node) GetColor() (result string) {
	if node == nil {
		result = "black"
	} else if node.Color == 1 {
		result = "black"
	} else {
		result = "red"
	}
	return
}

/*
* @receiver node
* @returns False case isnt the main root, true case it is
 */
func (node *Node) isMainRoot() (result bool) {
	if node.Parent == nil {
		result = true
	} else {
		result = false
	}
	return
}

/*
* @receiver node, node where we whish to flip a color
* @returns node with the color changed
 */
func (node *Node) flipColor() (result *Node) {
	if node.GetColor() == "black" {
		node.Color = 0
	} else {
		node.Color = 1
	}
	result = node
	return
}

/*
* @receiver node, the node where we want to make the flip of the colors happen
* @returns node with the colors fliped
 */
func (node *Node) ChangeColors() (result *Node) {
	node = node.flipColor()
	node.Left = node.Left.flipColor()
	node.Right = node.Right.flipColor()
	result = node
	return
}

/*
* @receiver node, node where we want to make a right rotation
* @returns node with the rotation left
 */
func (node *Node) RotateRight() (result *Node) {
	node = node.flipColor()
	temp := node.Left
	node.Left = nil
	temp.Right = node
	node = temp
	node.Parent = node.Right.Parent
	node.Right.Parent = node
	result = node
	return result
}

/*
* @receiver node, node which we want to rotate to the left
* @returns node with rotation left
 */
func (node *Node) RotateLeft() (result *Node) {
	node = node.flipColor()
	temp := node.Right
	temp.Left = node
	node.Right = nil
	node = temp
	node.Parent = node.Left.Parent
	node.Left.Parent = node
	result = node
	return
}

func (node *Node) GetChild(value int) (result *Node) {
	if value > node.Value {
		result = node.Right
	} else {
		result = node.Left
	}
	return
}

func (node *Node) GetUncle() (result *Node) {
	if node == nil {
		result = nil
		return
	}
	parent := node.Parent
	if parent == nil {
		result = nil
		return
	}
	grandParent := parent.Parent
	if grandParent == nil {
		result = nil
		return
	}
	if node.Value > grandParent.Value {
		result = grandParent.Left
	} else {
		result = grandParent.Right
	}
	return
}

/*
* @receiver node, the node where we want to insert a given value, if it is root or not depends of the lifecycle because of recursion
* @param value, This is the value we wish to insert, it will dictate how the node will make the insertion
* @returns a node that can either be the root or the last child inserted depending of the lifecycle
 */
func (node *Node) Insert(value int) (result *Node) {
	//? Case the coming node is null, we should retrieve the new node for insertion
	if node == nil {
		result = NewNode(value)
		return
	}
	/*
		? Part where we will iterate acording to the given value
	*/
	if node.Value > value {
		node.Left = node.Left.Insert(value)
		node.Left.Parent = node
	} else {
		node.Right = node.Right.Insert(value)
		node.Right.Parent = node
	}

	if node == nil {
		result = node
		return
	}

	if value > node.Value {
		if node.Left.GetColor() == "black" && node.Right.GetColor() == "red" {
			if value > node.Right.Value {
				if node.Right.Right.GetColor() == "red" {
					node = node.RotateLeft()
				}
			} else {
				if node.Right.Left.GetColor() == "red" {
					node.Right = node.Right.RotateRight()
					node = node.RotateLeft()
				}
			}
		}

		if node.Left.GetColor() == "red" && node.Right.GetColor() == "red" {
			if value > node.Right.Value {
				if node.Right.Right.GetColor() == "red" {
					node = node.ChangeColors()
				}
			} else {
				if node.Right.Left.GetColor() == "red" {
					node = node.ChangeColors()
				}
			}
		}
	} else {
		if node.Right.GetColor() == "black" && node.Left.GetColor() == "red" {
			if value < node.Left.Value {
				if node.Left.Left.GetColor() == "red" {
					node = node.RotateRight()
				}
			} else {
				if node.Left.Right.GetColor() == "red" {
					node.Left = node.Left.RotateLeft()
					node = node.RotateRight()
				}
			}
		}
		if node.Right.GetColor() == "red" && node.Right.GetColor() == "red" {
			if value > node.Right.Value {
				if node.Right.Right.GetColor() == "red" {
					node = node.ChangeColors()
				}
			} else {
				if node.Right.Left.GetColor() == "red" {
					node = node.ChangeColors()
				}
			}
		}
	}

	if node.isMainRoot() && node.GetColor() == "red" {
		node = node.flipColor()
	}

	result = node
	return
}

/*
* @receiver node
 */

func main() {
	root := &Node{Value: 5, Color: 1, Left: nil, Right: nil, Parent: nil}
	root = root.Insert(6)
	root = root.Insert(8)
	root = root.Insert(7)
	root = root.Insert(4)
	root = root.Insert(9)
	fmt.Println("Im the result ", root, root.Left, root.Left.Left, root.Right, root.Right.Left, root.Right.Right)
}
