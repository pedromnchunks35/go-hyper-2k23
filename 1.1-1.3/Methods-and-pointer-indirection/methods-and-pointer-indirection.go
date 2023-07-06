package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) calc(f int) {
	p.Age = p.Age + f
}

func calv2(p *Person, f int) {
	p.Age = p.Age + f
}

func main() {
	p := Person{Name: "Pedro", Age: 1}
	p.calc(2)
	fmt.Println(p.Age)
	calv2(&p, 2)
	fmt.Println(p.Age)
}
