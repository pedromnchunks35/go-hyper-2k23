package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot sqrt negative number %v", float64(e))
}
func Sqrt(x float64) (float64, error) {
	if x > 0 {
		return math.Sqrt(x), nil
	} else {
		return x, ErrNegativeSqrt(x)
	}
}

func main() {
	result, err := Sqrt(2)
	fmt.Println(result, err)
	result2, err2 := Sqrt(-2)
	fmt.Println(result2, err2)
}
