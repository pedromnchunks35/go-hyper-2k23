package main

import (
	"log"
	ch "own/chaincode"
	c "own/connection"
)

func main() {
	gateway, err := c.CreateConnection(
		"./msp/tlscacerts/ca.pem",
		"192.168.1.100:7051",
		"./msp/signcerts/cert.pem",
		"Org1MSP",
		"./msp/keystore/",
	)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer gateway.Close()
	channel := gateway.GetNetwork("channel1")
	contract := channel.GetContract("basic")
	ch.GetAssets(contract)
	ch.CreateAsset(contract, "6", "yellow", "1", "Carol", "10")
}
