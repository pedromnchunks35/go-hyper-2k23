package main

import "fmt"

func main() {
	i, j := 42, 2701
	p := &i                                              //point to i
	fmt.Println("Accessing value from the pointer ", *p) //read using the pointer
	*p = 21                                              // set i through the pointer
	fmt.Println("This is the value after setting it ", i)
	p = &j
	*p = *p / 37                                                     //divide j through the pointer
	fmt.Println("Value after address change and after division ", j) // see the new value of j
}
