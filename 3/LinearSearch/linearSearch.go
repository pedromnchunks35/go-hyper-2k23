package main

import "fmt"

func linearsearch(data *[]int, key int) bool {
	for _, item := range *data {
		if item == key {
			return true
		}
	}
	return false
}

func main() {
	items := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, -2}
	fmt.Println(linearsearch(&items, -2))
}
