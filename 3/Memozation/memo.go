package main

import (
	"fmt"
)

type MemoizedFunction struct {
	memo map[int]int
}

func NewMemoizedFunction() *MemoizedFunction {
	return &MemoizedFunction{
		memo: make(map[int]int),
	}
}

func (m *MemoizedFunction) Fibonnaci(n int) int {
	if result, ok := m.memo[n]; ok {
		fmt.Println("Reused result")
		return result
	}
	if n <= 1 {
		m.memo[n] = n
		return n
	}
	result := m.Fibonnaci(n-1) + m.Fibonnaci(n-2)
	m.memo[n] = result
	return result
}

func main() {
	memoFunc := NewMemoizedFunction()
	fmt.Println(memoFunc.Fibonnaci(10))
	fmt.Println(memoFunc.Fibonnaci(10))
	fmt.Println(memoFunc.Fibonnaci(5))
	fmt.Println(memoFunc.Fibonnaci(5))
}
