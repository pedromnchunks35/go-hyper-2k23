package main

import (
	"fmt"
	"io"
)

type MyReader struct{}

func (r MyReader) Read(p []byte) (n int, err error) {
	for i := 0; i < len(p); i++ {
		p[i] = 'A'
	}
	return len(p), nil
}

func main() {
	reader := MyReader{}
	buffer := make([]byte, 10)
	for {
		_, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		fmt.Println(string(buffer))
	}
}
