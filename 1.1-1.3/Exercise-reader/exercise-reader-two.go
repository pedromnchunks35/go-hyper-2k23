package main

import (
	"fmt"
	"io"
)

type MyReader struct{}

func (m MyReader) Read(buffer []byte) error {
	for i := 0; i < len(buffer); i++ {
		buffer[i] = 'A'
	}
	return nil
}

func main() {
	reader := MyReader{}
	buffer := make([]byte, 10)
	for {
		err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		fmt.Println(string(buffer))
	}

}
