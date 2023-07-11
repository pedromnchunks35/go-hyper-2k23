package main

import "fmt"

func InsertionSort(items *[]int) {
	n := len(*items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if (*items)[j-1] > (*items)[j] {
				(*items)[j-1], (*items)[j] = (*items)[j], (*items)[j-1]
			}
			j -= 1
		}
	}
}

func main() {
	array := []int{10, -2, 3, 30, 40, 50, 2, 3, 1, 4, 7, -20}
	fmt.Printf("Array before: %v\n", array)
	InsertionSort(&array)
	fmt.Printf("Array After: %v\n", array)
}
