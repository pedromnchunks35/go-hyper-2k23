package main

import (
	"fmt"
)

type Peps int

func (p Peps) printSomething() {
	fmt.Println(p)
}

func main() {
	r := Peps(2)
	r.printSomething()
}
