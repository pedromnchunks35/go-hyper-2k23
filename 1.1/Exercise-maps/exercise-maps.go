package main

import (
	"fmt"
	"strings"
	"unicode"
)

func WordCount(s string) map[string]int {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	separedwords := strings.FieldsFunc(s, f)
	var m map[string]int = make(map[string]int)
	for i := 0; i < len(separedwords); i++ {
		m[separedwords[i]] = 1
	}
	return m
}

func main() {
	result := WordCount("Hello world im here")
	fmt.Println(result)
}
