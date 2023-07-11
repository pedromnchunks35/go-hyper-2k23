package main

import "fmt"

func BubbleSort(sliceToSort *[]int) {
	for i := 0; i < len(*sliceToSort)-1; i++ {
		for j := 0; j < len(*sliceToSort)-i-1; j++ {
			if (*sliceToSort)[j] > (*sliceToSort)[j+1] {
				//? Notice that normally we would store the value in temporary fields because by changing the value of the first, when assigning that value to the secound we would end up
				//? With a slice with duplicated members
				//? But by assigning this way, go does that alone
				(*sliceToSort)[j], (*sliceToSort)[j+1] = (*sliceToSort)[j+1], (*sliceToSort)[j]
			}
		}
	}
}
func main() {
	toOrder := []int{4, 6, 7, 8, 9, 1, 2, 3, 10, 12, 14, 16}
	fmt.Printf("The array before: %v \n", toOrder)
	BubbleSort(&toOrder)
	fmt.Printf("The array after: %v \n", toOrder)
}
