package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice := []int{3, 1, 2, 4, 5, 6}
	quicksort(slice, &slice)
}

func quicksort(a []int, main *[]int) []int {
	fmt.Println("------START ITERACTION----")
	fmt.Println("Array to iteract: ", a)
	if len(a) < 2 {
		return a
	}
	fmt.Printf("Main array 1º: %v\n", *main)
	left, right := 0, len(a)-1
	fmt.Println("left: ", left)
	fmt.Println("right: ", right)
	pivot := rand.Int() % len(a)
	fmt.Println("This is the Pivot: ", pivot)
	a[pivot], a[right] = a[right], a[pivot]
	fmt.Println("pivot<>right: ", a)
	fmt.Printf("Main array 2º: %v\n", *main)
	for i, _ := range a {
		fmt.Println("This is the pos ", i)
		if a[i] < a[right] {
			fmt.Printf("left: %v\n", left)
			fmt.Printf("i: %v\n", i)
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	fmt.Printf("Main array 4º: %v \n", *main)
	a[left], a[right] = a[right], a[left]
	fmt.Println("left<>right: ", a)
	fmt.Printf("Main array 5º: %v \n", *main)
	fmt.Println("---Next Iteraction---")

	quicksort(a[:left], main)
	quicksort(a[left+1:], main)

	return a
}
