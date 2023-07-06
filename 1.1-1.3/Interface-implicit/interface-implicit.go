package main

import "fmt"

type I interface {
	Print()
}

type Person struct {
	Name string
}

func (p Person) Print() {
	fmt.Println(p.Name)
}

func main() {
	var inf I = Person{Name: "Peter"}
	inf.Print()
}
