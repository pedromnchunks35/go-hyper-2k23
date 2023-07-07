package main

import (
	"flag"
	"fmt"
)

type Error struct {
	Message string
}

func (e Error) Error() string {
	return e.Message
}

func Read(input string, position int) (newpos int, buffer []byte, err Error) {
	buffer = make([]byte, 10)
	for i := 0; i < len(buffer); i++ {
		if position == len(input) {
			return 0, buffer, Error{Message: "EOF"}
		}
		buffer[i] = input[position]
		position++
	}
	return position, buffer, Error{}
}

func main() {
	test := flag.String("i", "", "This is the input we want to put inside of a buffer")
	flag.Parse()
	value := *test
	if value == "" {
		message := Error{Message: "Sorry but you need to provide something in the input"}
		fmt.Println(message)
		return
	}
	position := 0
	for {
		newpos, buffer, err := Read(value, position)
		position = newpos
		fmt.Printf("%v", string(buffer))
		if err.Message != "" {
			fmt.Println("\n ", err)
			return
		}
	}
}
