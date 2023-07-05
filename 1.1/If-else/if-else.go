package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else if k := math.Pow(x, n); k < lim+10 {
		return v
	} else {
		fmt.Printf("No ifs worked \n")
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
		pow(3, 5, 10),
	)
}
