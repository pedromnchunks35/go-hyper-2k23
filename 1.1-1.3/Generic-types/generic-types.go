package main

import "fmt"

type List[T any] struct {
	next *List[T]
	val  T
}

type Test[X, K any] struct {
	val    X
	valtwo K
}

func main() {
	x := List[string]{next: nil, val: "head"}
	v := List[string]{next: &x, val: "hello"}
	fmt.Println(x, v)
	b := Test[string, int]{"ALo", 2}
	fmt.Println(b)
}
