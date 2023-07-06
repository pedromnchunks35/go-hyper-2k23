package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

func (_ip IPAddr) String() string {
	var result string
	for i := 0; i < len(_ip); i++ {
		result += strconv.Itoa(int(_ip[i])) + "."
	}
	return result[:len(result)-1]
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v \n", name, ip)
	}
}
