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

func (head Tree) New(numbers *[]int) {
	randomNum := rand.Intn(len(*numbers))
	fmt.Println(randomNum)
}

func removeByIndex(index int, numbers *[]int) {
	n := *numbers
	*numbers = append(n[:index], n[index+1:]...)
}

func main() {
	//leafs := 0
	k := 1
	var numbers []int
	for i := 1; i < 11; i++ {
		numbers = append(numbers, i*k)
	}
	tree := Tree{Left: &Tree{Left: nil, Value: 1, Right: nil}}
	tree.New(&numbers)
}
