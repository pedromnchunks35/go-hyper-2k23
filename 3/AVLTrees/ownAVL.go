package main

import "fmt"

type Node struct {
	Heigth int
	Value  int
	Left   *Node
	Right  *Node
}

/*
* @receiver node, The node where we want to extract the heigth
* @returns 0 case the reference is nil, otherwise it returns the heigth associated to that node
 */
func (node *Node) GetHeigth() (result int) {
	if node == nil {
		result = 0
	} else {
		result = node.Heigth
	}
	return
}

/*
* @param x, one of the values to compare
* @param y, another value to compare
* @returns the biggest value between two given values
 */
func max(x, y int) (result int) {
	if x > y {
		result = x
	} else {
		result = y
	}
	return
}

/*
* @receiver node, the node where we want to check out the balance factor that he has to offer
* @returns the balance factor, which is calculated based on the formula BF= L - R, where L is
* the left side height and R is the Right side heigth
 */
func (node *Node) BalanceFactor() (result int) {
	return node.Left.GetHeigth() - node.Right.GetHeigth()
}

/*
* @param value, the value that the new node must have
* @returns a new node
 */
func NewNode(value int) (result *Node) {
	result = &Node{Value: value, Heigth: 1, Left: nil, Right: nil}
	return
}

/*
* @param node, the node we want to make the childrens to rotate to the left
* @returns node, after the rotation to the left
 */
func RotateLeft(node *Node) (result *Node) {
	tempRight := node.Right
	node.Right = nil
	tempRight.Left = node
	node = tempRight
	node.Left.Heigth = max(node.Left.Left.GetHeigth(), node.Left.Right.GetHeigth()) + 1
	node.Heigth = max(node.Left.GetHeigth(), node.Right.GetHeigth()) + 1
	result = node
	return
}

/*
* @param node, the node we want to make the childrens to rotate to the right
* @returns node after the rotation to the Right
 */
func RotateRight(node *Node) (result *Node) {
	tempLeft := node.Left
	node.Left = nil
	tempLeft.Right = node
	node = tempLeft
	node.Right.Heigth = max(node.Right.Left.GetHeigth(), node.Right.Right.GetHeigth()) + 1
	node.Heigth = max(node.Left.GetHeigth(), node.Right.GetHeigth()) + 1
	result = node
	return
}

/*
* @param node, the node where we want to insert the value
* @param value, the value that we are assigning to the new node we wish to create
* @returns A new node or the result from the rotation, depending of the lifecycle of the function
 */
func InsertNode(node *Node, value int) (result *Node) {
	/*
		? Stop condition for when we get to the final of the tree

		? The value returned from here will be used more ahead to either be apart of the left or right
		? side of the node we are now, which is the result from the recursive operations we are making
	*/
	if node == nil {
		result = NewNode(value)
		return
	}
	/*
		? This part is responsible of navigating throught the tree using recursion, it will stop the
		? navigation when reaching the stopping point that we mention before
	*/
	if value > node.Value {
		node.Right = InsertNode(node.Right, value)
	} else {
		node.Left = InsertNode(node.Left, value)
	}
	/*
		? Part for calculating the Height of the node we current are

		? Note that this part will only take place after inserting the new node
		? which is when it reaches the stop point in the beggining of the function and assigns that
		? new node to either the left or right, previous mentioned

		? Note that the Height is always the max of Left and Right Heights plus 1, in order to not
		? loose the track
	*/
	node.Heigth = max(node.Left.GetHeigth(), node.Right.GetHeigth()) + 1
	/*
		? Part where we see for where we should rotate

		? To decide if we should rotate or not we should check the balanceFactor function

		? Case the balanceFactor is bigger than 1, we rotate Right because we are on the Left
		? , case the balanceFactor is lesser than -1, we rotate Left because we are on the right

		? We should note that if we have balance factor bigger than 1 and the just inserted node is
		? bigger than the left value, we should rotate to the left and assign that value to the left of the
		? node, which will create a diagonal result for the left. After that we rotate the current node to the right

		? The same goes for the lesser that -1, the diagonal is for the right, and the directions are also different
		? but the idea is the same
	*/
	balance := node.BalanceFactor()
	if balance > 1 {
		if node.Left.Value < value {
			node.Left = RotateLeft(node.Left)
			result = RotateRight(node)
			return
		} else {
			result = RotateRight(node)
			return
		}
	}

	if balance < -1 {
		if node.Right.Value > value {
			node.Right = RotateRight(node.Right)
			result = RotateLeft(node)
			return
		} else {
			result = RotateLeft(node)
			return
		}
	}
	//? After all done, we return the root simply
	result = node
	return
}

func main() {
	root := InsertNode(nil, 2)
	root = InsertNode(root, 4)
	root = InsertNode(root, 3)
	root = InsertNode(root, 5)
	root = InsertNode(root, 6)
	fmt.Println(root)
}
