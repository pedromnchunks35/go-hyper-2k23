package main

import "fmt"

func Pic(dx, dy int) [][]uint8 {
	var result [][]uint8
	result = make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		result[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			result[i][j] = uint8((i * j) / 2)
		}
	}
	return result
}

func PrintArray(result [][]uint8) {
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			fmt.Printf("--%d--", result[i][j])
		}
	}
}

func main() {
	result := Pic(10, 20)
	PrintArray(result)
}
