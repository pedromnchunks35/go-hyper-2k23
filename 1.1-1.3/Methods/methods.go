package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) GetName() {
	fmt.Println(p.Name)
}

func main() {
	person := Person{Name: "Pedro", Age: 2}
	person.GetName()
}
