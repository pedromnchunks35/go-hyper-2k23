package main

import "fmt"

func fibonnaci() func() int {
	//INIT VARIABLES
	a, b := 0, 1
	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}

func main() {
	f := fibonnaci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
