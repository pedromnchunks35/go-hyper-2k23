package main

import (
	"fmt"
	"time"
)

func loopas(c chan int, k chan int) {
	for {
		select {
		case <-c:
			fmt.Println("C is here")
		case <-k:
			fmt.Println("k is here")
		default:
			fmt.Println("It is a new world")
			time.Sleep(10 * time.Second)
		}
	}
}

func main() {
	c := make(chan int)
	k := make(chan int)
	go func() {
		for {
			c <- 1
			time.Sleep(2 * time.Second)
			k <- 2
		}
	}()
	loopas(c, k)
}
