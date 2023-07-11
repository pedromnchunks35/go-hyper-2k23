package main

import "fmt"

func SelectionSort(slice *[]int) {
	var minIndex int
	for i := 0; i < len(*slice)-1; i++ {
		minIndex = i
		for j := i + 1; j < len(*slice); j++ {
			if (*slice)[j] < (*slice)[minIndex] {
				minIndex = j
			}
		}
		(*slice)[i], (*slice)[minIndex] = (*slice)[minIndex], (*slice)[i]
	}
}
func main() {
	array := []int{20, 30, 22, 1, 2, 3, 5, 8, 0, 6, 100}
	fmt.Printf("The slice before: %v\n", array)
	SelectionSort(&array)
	fmt.Printf("The slice after: %v\n", array)
}
