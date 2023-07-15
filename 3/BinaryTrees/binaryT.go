package main

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func removeNumberFromArray(numbers *[]int, pos int) {
	(*numbers) = append((*numbers)[:pos], (*numbers)[pos+1:]...)
}

func insertLeaf(leaf *Tree, num int) {
	if leaf.Value < num {
		if leaf.Right != nil {
			leaf = leaf.Right
			insertLeaf(leaf, num)
		} else {
			leaf.Right = &Tree{Value: num, Left: nil, Right: nil}
		}
	} else {
		if leaf.Left != nil {
			leaf = leaf.Left
			insertLeaf(leaf, num)
		} else {
			leaf.Left = &Tree{Value: num, Left: nil, Right: nil}
		}
	}
}

func search(l *Tree) {
	if l != nil {
		fmt.Println("Value of the tree ", l.Value)
		search(l.Left)
		search(l.Right)
	}
}

func searchPost(l *Tree) {
	if l != nil {
		search(l.Left)
		search(l.Right)
		fmt.Println("Value of the tree ", l.Value)
	}
}

func (t *Tree) InsertLeafs(numbers *[]int) {
	//? Inserting first leaf
	randomPos := rand.Intn(len(*numbers))
	newValue := (*numbers)[randomPos]
	currentLeaf := t
	removeNumberFromArray(numbers, randomPos)
	t.Value = newValue
	t.Left = nil
	t.Right = nil
	for len(*numbers) != 0 {
		currentLeaf = t
		randomPos = rand.Intn(len(*numbers))
		newValue = (*numbers)[randomPos]
		removeNumberFromArray(numbers, randomPos)
		insertLeaf(currentLeaf, newValue)
	}
}

func main() {
	tree := Tree{}
	numbers := []int{30, 20, 10, 2, -2, 1, 3, 4, 5}
	tree.InsertLeafs(&numbers)
	search(&tree)
}
