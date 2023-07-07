package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("Im like a file you know?")
	buffer := make([]byte, 10)
	for {
		_, err := reader.Read(buffer)
		if err == io.EOF {
			fmt.Println("\nReached the end")
			break
		}
		fmt.Printf("%v", string(buffer))
		buffer = make([]byte, 10)
	}
}
