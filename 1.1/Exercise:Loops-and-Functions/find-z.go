/*
As a way to play with functions and loops,
let's implement a square root function: given a number x,
we want to find the number z for which z² is most nearly x.
*/
package main

import "fmt"

func sqrt(_x float64) float64 {
	return _x * _x
}

func eval(_x, _z float64) float64 {
	result := (sqrt(_z) - _x)
	if result < 0 {
		return result * -1
	}
	return result
}

func demo(_x float64) {
	var result float64 = 1000
	var final float64
	for i := 1.0; i < 20.0 && result > 1; i++ {
		if val := eval(_x, i); val < result {
			result = val
			final = i
			fmt.Printf("The closest number so far is: %v \n", result)
		} else {
			fmt.Printf("Not that close \n")
		}
	}
	fmt.Printf("The best variation is %v. The best value is %v \n", result, final)
}

func main() {
	demo(40)
}
