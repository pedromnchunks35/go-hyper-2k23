package main

import "fmt"

type Node struct {
	Value       int
	Color       int
	Left, Right *Node
	Parent      *Node
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
	if node.Color == 1 {
		result = "black"
	} else {
		result = "red"
	}
	return
}

func (node *Node) GetParent() (result *Node) {
	result = node.Parent
	return
}

func (node *Node) GetGrandParent() (result *Node) {
	result = node.GetParent().GetParent()
	return
}

/*
* @receiver parent, node where we want to get the brother
* @returns the brother of the current node
 */
func (node *Node) GetUncle() (result *Node) {
	grandParent := node.GetGrandParent()
	if node.GetParent().GetValue() > grandParent.GetValue() {
		result = grandParent.Left
	} else {
		result = grandParent.Right
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
	//? PUT THE GRANDFATHER AS TEMP VAR
	temp := node
	//? THE GRANDFATHER NOW IS THE LEFT OF THE GRANDFATHER
	node = node.Left
	//? GRANDFATHER THAT NOW IS THE LEFT NEEDS TO HAVE THE GRANDFATHER PARENT
	node.Parent = temp.Parent
	//? OLD GRANDFATHER NOW IS THE RIGHT OF THE NEW GRANDFATHER
	node.Right = temp
	//? THE PARENT OF THE OLD GRANDFATHER NOW IS THE NEW GRANDFATHER
	node.Right.Parent = node
	//? Make the temp left nil
	node.Right.Left = nil
	result = node
	return
}

/*
* @receiver node, node which we want to rotate to the left
* @returns node with rotation left
 */
func (node *Node) RotateLeft() (result *Node) {
	//? PUT THE GRANDFATHER AS TEMP VAR
	temp := node
	//? THE GRANDFATHER NOW IS THE Right OF THE GRANDFATHER
	node = node.Right
	//? GRANDFATHER THAT NOW IS THE RIGHT NEEDS TO HAVE THE GRANDFATHER PARENT
	node.Parent = temp.Parent
	//? THE LEFT OF THE NEW GRANDFATHER IS THE OLD GRANDFATHER
	node.Left = temp
	//? THE PARENT OF THE OLD GRANDFATHER NOW IS THE NEW GRANDFATHER
	node.Left.Parent = node
	//? REMOVE THE RIGHT FROM THE NEW LEFT
	node.Left.Right = nil
	result = node
	return
}

/*
* @receiver node, the node which we want to verify
* @returns node with enought modifications
 */
func (node *Node) VerifyNode() (result *Node) {
	//? If the child or the parent are the main root, return node
	if node.isMainRoot() || node.GetParent().isMainRoot() {
		result = node
		return
	}
	//? Start normal operation
	grandParent := node.GetGrandParent()
	child := node
	father := node.GetParent()
	uncle := node.GetUncle()
	/*
		? In this part we verify the pattern of child ,f ather red and uncle nil(black) or black

		? If this happens we need to apply two possible scenarios: rotate right, rotate left right, rotate left
		? or rotate right left
	*/
	if child.GetColor() == "red" && father.GetColor() == "red" && (uncle == nil || uncle.GetColor() == "black") {
		//? Rules respected to the left side
		if grandParent.GetValue() > father.GetValue() {
			if father.GetValue() > child.GetValue() {
				grandParent = grandParent.RotateRight()
				child = grandParent.Left.Left
				father = grandParent.Left
			} else {
				father = father.RotateLeft()
				child = father.Right
				grandParent = father.Parent
				grandParent = grandParent.RotateRight()
			}
			//? Rules for the right side
		} else {
			if father.GetValue() < child.GetValue() {
				grandParent = father.RotateLeft()
				child = grandParent.Right.Right
				father = grandParent.Right
			} else {
				father = father.RotateRight()
				child = father.Right
				grandParent = father.Parent
				grandParent = grandParent.RotateLeft()
			}
		}
	}
	/*
		? This is the condition where child,father and uncle are red, this means we should
		? change all the colors
	*/
	if uncle != nil && child.GetColor() == "red" && father.GetColor() == "red" && uncle.GetColor() == "red" {
		grandParent = grandParent.ChangeColors()
	}

	if grandParent.isMainRoot() {
		if grandParent.GetColor() == "red" {
			grandParent = grandParent.flipColor()
		}
	}
	//? Check if the grandParent is the main Root, if it is, make sure it becomes black
	result = child
	return
}

func main() {

	three := &Node{Value: 3, Left: nil, Right: nil, Color: 0, Parent: nil}
	five := &Node{Value: 5, Left: three, Right: nil, Color: 0, Parent: nil}
	three.Parent = five
	root := &Node{Value: 6, Left: five, Right: nil, Color: 1, Parent: nil}
	five.Parent = root
	fmt.Printf("%p \n", root)
	fmt.Printf("%p \n", three)
	root = three.VerifyNode()
	fmt.Printf("%p \n", root)
	fmt.Println(three.Parent)
}
