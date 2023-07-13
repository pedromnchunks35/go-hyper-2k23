package main

import "fmt"

func main(){
x:=make(map[int][]int)
x[1] = append(x[1],3)
x[1] = append(x[1],4)
fmt.Println(x)
}
