package main

import (
	"fmt"
)

type Reader struct{}

func (r Reader) Error() string {
	return "EOF verified"
}

func (r Reader) Read(position int, input string) (length int, buffer []byte, err error) {
	buffer = make([]byte, 10)
	for i := 0; i < len(buffer); i++ {
		if position == len(input) {
			return 0, nil, r
		}
		buffer[i] = input[position]
		position++
	}
	return position, buffer, nil
}

func main() {
	input := "Im just a self made input, im op op op hahahahahahahhahaha"
	reader := Reader{}
	position := 0
	for {
		newpos, buffer, err := reader.Read(position, input)
		if err != nil {
			fmt.Println("\n ", err)
			break
		}
		fmt.Printf("%v", string(buffer))
		position = newpos
	}
}
