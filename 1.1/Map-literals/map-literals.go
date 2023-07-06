package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

var m = map[string]Person{
	"Pedro": Person{
		Name: "Pedro",
		Age:  12,
	},
	"Elsa": Person{
		Name: "Elsa",
		Age:  22,
	},
}

func main() {
	fmt.Println(m)
}
