package main

import "fmt"

type test struct {
	Value     int
	Position  *test
	Position2 *test
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
	t2 := &test{Value: 3, Position: nil}
	t := &test{Value: 2, Position: t2}
	fmt.Printf("{Pos:%p,Value:%v},{Pos:%p,Value:%v}\n", t, t.Value, t2, t2.Value)
	t = t.tradeof()
	t2 = t.Position
	fmt.Printf("{Pos:%p,Value:%v},{Pos:%p,Value:%v},{Pos: %p,Value: %v}\n", t, t.Value, t2, t2.Value, t.Position, t.Position.Value)
	t.changeValue()
	fmt.Printf("{Pos:%p,Value:%v},{Pos:%p,Value:%v},{Pos: %p,Value: %v}\n", t, t.Value, t2, t2.Value, t.Position, t.Position.Value)
}
