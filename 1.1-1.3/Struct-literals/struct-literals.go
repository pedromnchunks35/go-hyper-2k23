package main

import "fmt"

type Vertex struct {
	Name string
	Age  int
}

var (
	v1 = Vertex{
		Name: "Pedro",
		Age:  11,
	}
)

func main() {
	fmt.Println(v1)
}
