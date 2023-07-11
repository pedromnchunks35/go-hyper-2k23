package main

import (
"fmt"
"math/rand"
)

func main(){
fmt.Println(rand.Int()%3)
x:=[]int{1,2,3,4,5}
fmt.Println(x[1:])
fmt.Println(x[:2])
}
