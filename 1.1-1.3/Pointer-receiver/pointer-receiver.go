package main

import "fmt"

type Person struct {
	Name      string
	Age       int
	randomVal int
}

func (p *Person) changeName(name string) {
	p.Name = name
}

func main() {
	var p Person = Person{Name: "Pedro", Age: 2, randomVal: 3}
	fmt.Println(p.Name)
	p.changeName("Elsa")
	fmt.Println(p.Name)
}
