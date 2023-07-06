package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

var m map[string]Person

func main() {
	m = make(map[string]Person)
	m["test"] = Person{
		Name: "Pedro",
		Age:  12,
	}
	fmt.Println(m["test"])
}
