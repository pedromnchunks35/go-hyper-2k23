package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) String() string {
	return "THis is the output whatever the value of name"
}

func main() {
	t := Person{Name: "P"}
	fmt.Println(t)
}
