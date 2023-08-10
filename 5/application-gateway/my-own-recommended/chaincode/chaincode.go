package chaincode

import (
	"fmt"
	"log"
	"own/utils"

	"github.com/hyperledger/fabric-gateway/pkg/client"
)

func GetAssets(contract *client.Contract) {
	res, err := contract.Evaluate("GetAllAssets")
	if err != nil {
		log.Fatalf("error evaluatiating assets %v", err)
	}
	string, err := utils.FormatJSON(res)
	if err != nil {
		log.Fatalf("error formating to json %v", err)
	}
	fmt.Println(string)
}
