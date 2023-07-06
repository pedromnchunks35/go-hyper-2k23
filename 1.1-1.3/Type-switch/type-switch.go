package main

import "fmt"

func do(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("It is a int")
	case string:
		fmt.Println("It is a string")
	default:
		fmt.Println("It is something we dont know about")
	}
}

func main() {
	do(21)
}
