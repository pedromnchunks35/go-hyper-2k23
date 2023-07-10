package main

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

/*
* @param (receiver) head, this is the pointer to the head of the tree, the starting point
* @param numbers, this is the pointer to the slice of numbers we must put in the binary tree
* @returns nothing, but it does start the insertion of the numbers in the binary tree
 */
func (head *Tree) New(numbers *[]int) {
	for {
		var path []int
		tail := head
		if len((*numbers)) == 0 {
			break
		}
		randomIndex := rand.Intn(len(*numbers))
		NewLeaf(tail, (*numbers)[randomIndex], path)
		removeByIndex(randomIndex, numbers)
	}
}

/*
* @param t, a member of a tree, the first iteration is the head but it mutates by iteration
* @param numero, this is the number that we want to put in the binary tree in a ordered way
* @param path, this is the array that tells us the path to reach the number we just inserted (0=left,1=right)
* @returns nothing but adds new leafs to the tree
 */
func NewLeaf(t *Tree, numero int, path []int) {
	if numero > t.Value {
		path = append(path, 1)
		if t.Right != nil {
			NewLeaf(t.Right, numero, path)
		} else {
			t.Right = &Tree{Left: nil, Value: numero, Right: nil}
			fmt.Printf("The %v is placed. The path is %v \n", numero, path)
		}
	} else {
		path = append(path, 0)
		if t.Left != nil {
			NewLeaf(t.Left, numero, path)
		} else {
			t.Left = &Tree{Left: nil, Value: numero, Right: nil}
			fmt.Printf("The %v is placed. The path is %v \n", numero, path)
		}
	}
}

/*
* @param index, index that we want to remove
* @param numbers, is the pointer for the slice where we want to remove a certain member
* @returns, nothing but it removes the desired index from the slice we are pointing to
 */
func removeByIndex(index int, numbers *[]int) {
	n := *numbers
	*numbers = append(n[:index], n[index+1:]...)
}

/*
* @param t, head of a certain tree
* @param ch, the channel for where we will send the leaf values
* @returns nothing, but this is the start of the iteration all over the leafs
 */
func Walk(t *Tree, ch chan int) {
	walkRecursive(t, ch)
}

/*
* @param t, leaf from a certain tree
* @param ch, the channel for where we will send the leaf values
* @returns nothing but it will iterate all over all branches if the leaf exists
 */
func walkRecursive(t *Tree, ch chan int) {
	if t == nil {
		return
	}
	ch <- t.Value
	walkRecursive(t.Left, ch)
	walkRecursive(t.Right, ch)
}

/*
* @param t, the head of a tree
* @param t2, another head of a different tree
* @returns, either if t and t2 are the same by comparing leafs using channels
 */
func Same(t *Tree, t2 *Tree) bool {
	isEqual := true
	channel := make(chan int)
	channelTwo := make(chan int)
	go Walk(t, channel)
	go Walk(t2, channelTwo)
	for i := 0; i < 10; i++ {
		x, y := <-channel, <-channelTwo
		if isEqual == true {
			isEqual = (x == y)
		}
	}
	return isEqual
}

func main() {
	//? GENERATE THE NUMBERS
	k := 1
	var numbers []int
	var numbers2 []int
	for i := 1; i < 11; i++ {
		numbers = append(numbers, i*k)
		numbers2 = append(numbers2, i*k)
	}
	//? GENERATE FIRST TREE
	randomNum := rand.Intn(len(numbers))
	fmt.Printf("The %v is placed. \n", (numbers)[randomNum])
	tree := Tree{Left: nil, Value: (numbers)[randomNum], Right: nil}
	removeByIndex(randomNum, &numbers)
	tree.New(&numbers)
	//? GENERATE THE SECOUND TREE
	randomNum2 := rand.Intn(len(numbers2))
	fmt.Printf("The %v is placed. \n", (numbers2)[randomNum2])
	tree2 := Tree{Left: nil, Value: (numbers2)[randomNum2], Right: nil}
	removeByIndex(randomNum2, &numbers2)
	tree2.New(&numbers2)
	//? VERIFY IF THE SAME TREE IS EQUAL TO THE SAME TREE
	isSame := Same(&tree, &tree)
	fmt.Printf("The tree is the same (checking the same tree against the same tree)? %v \n", isSame)
	//? VERIFY IF DIFFERENT TREES ARE QUAL
	isSame2 := Same(&tree, &tree2)
	fmt.Printf("The tree is the same (two different trees were generated here)? %v \n", isSame2)
}
