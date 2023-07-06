package main

import "fmt"

func add(x int, y int) int {
	return x + y
}
func main() {
	fmt.Println("This is the result ", add(2, 3))
}
