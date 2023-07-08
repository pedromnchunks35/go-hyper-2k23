package main

import "fmt"

func justSelect(c chan int, k chan int) {
	for {
		select {
		case <-c:
			fmt.Println("Im c")
		case <-k:
			fmt.Println("Im k")
			return
		}
	}
}

func main() {
	c := make(chan int)
	k := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		k <- 0
	}()
	justSelect(c, k)
}
