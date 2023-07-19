package main

import "fmt"

type test struct {
	Value     int
	Position  *test
	Position2 *test
}

func (t *test) changePointers() {
	h := &test{Value: 2, Position: nil}
	t = h
}

func (t *test) changeValue() {
	temp := (*t).Value
	t.Value = (*t).Position.Value
	t.Position.Value = temp
}

func (t *test) tradeof() (result *test) {
	temp := t
	fmt.Printf("This is temp %p\n", temp)
	t = t.Position
	t.Position = temp
	fmt.Printf("This is t.position after giving temp %p\n", t.Position)
	result = t
	return
}

func main() {
	t := &test{Value: 1}
	fmt.Printf("%p %v\n", t, t)
	t.changePointers()
	fmt.Printf("%p %v\n", t, t)
}
