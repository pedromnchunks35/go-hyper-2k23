package main

import "fmt"

func MergeSort(items []int) []int {
	var length = len(items)
	if length == 1 {
		return items
	}
	middle := int(length / 2)
	left := make([]int, middle)
	right := make([]int, length-middle)
	for i := 0; i < length; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}
	return Merge(MergeSort(left), MergeSort(right))
}

func Merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))
	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}
	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}
	return
}

func main() {
	array := []int{-20, 30, 2, 1, 4, 5, 200, 1000, -5}
	fmt.Printf("This is the array before: %v\n", array)
	result := MergeSort(array)
	fmt.Printf("This is the new ordered array: %v \n", result)
}
