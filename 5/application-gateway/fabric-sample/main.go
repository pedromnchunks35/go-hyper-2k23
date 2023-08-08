/*
Copyright 2020 IBM All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	//? Just for us to know, this isnt the config.yaml file, it is a config binarie
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	//? The gateway operations of fabric-sdk-go
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

func main() {
	log.Println("============ application-golang starts ============")
	//? This sets a enviromental var in the system we are wroking
	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	if err != nil {
		log.Fatalf("Error setting DISCOVERY_AS_LOCALHOST environment variable: %v", err)
	}
	//? Establish the walletpath
	walletPath := "wallet"
	//? Delete every files inside of the wallet to create a new wallet
	os.RemoveAll(walletPath)
	/*
		? A wallet is a struct where we store every information about the identity to connecto the hyper ledger networker
		? In this part of the file we create a instance of the wallet held in memory
	*/
	wallet, err := gateway.NewFileSystemWallet(walletPath)
	if err != nil {
		log.Fatalf("Failed to create wallet: %v", err)
	}
	/*
		? Check if a wallet with the user "appUser" exists
	*/
	if !wallet.Exists("appUser") {
		/*
			? Case it does not exist, we populate the wallet (this is a func that is present in this file itself)
		*/
		err = populateWallet(wallet)
		if err != nil {
			log.Fatalf("Failed to populate wallet contents: %v", err)
		}
	}
	//? Get the common connection profile file, this is the description of the connection
	//? We are trying to achieve
	ccpPath := filepath.Join(
		"..",
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"connection-org1.yaml",
	)
	/*
		? This is where we will make the connection to our fabric network
		? We need a common connection profile file, which is a file that checks the connection
		? We need the identity that we just created in our wallet
		? We can also add new options which we only have the timout
	*/
	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, "appUser"),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gateway: %v", err)
	}
	//? Close it after everything
	defer gw.Close()
	/*
		? Grab the channel name (default name is "mychannel")
		? If the name is another the enviroment variable "CHANNEL_NAME", must point to another name
	*/
	channelName := "mychannel"
	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}

	log.Println("--> Connecting to channel", channelName)
	//? Get the channel
	network, err := gw.GetNetwork(channelName)
	if err != nil {
		log.Fatalf("Failed to get network: %v", err)
	}
	/*
		? Now Grab the chaincode name, the default is basic
		? If the name is another the enviroment variable "CHAINCODE_NAME" must be different
	*/
	chaincodeName := "basic"
	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}
	//? it is throught the chaincode name that we can grab the contract using the "network"(channel)
	log.Println("--> Using chaincode", chaincodeName)
	contract := network.GetContract(chaincodeName)
	/*
		?INIT THE LEDGER HERE
	*/
	log.Println("--> Submit Transaction: InitLedger, function creates the initial set of assets on the ledger")
	result, err := contract.SubmitTransaction("InitLedger")
	if err != nil {
		log.Fatalf("Failed to Submit transaction: %v", err)
	}
	log.Println(string(result))

	log.Println("--> Evaluate Transaction: GetAllAssets, function returns all the current assets on the ledger")
	/*
		? INVOKE THE METHOD GetAllAssets (if it is a get we use the "Evaluate" funciton)
	*/
	result, err = contract.EvaluateTransaction("GetAllAssets")
	if err != nil {
		log.Fatalf("Failed to evaluate transaction: %v", err)
	}
	log.Println(string(result))

	log.Println("--> Submit Transaction: CreateAsset, creates new asset with ID, color, owner, size, and appraisedValue arguments")
	/*
		? When we need to write something into the smart contract we need to use the SubmitTransaction, note that firstly it comes the name of the function and latter the args
	*/
	result, err = contract.SubmitTransaction("CreateAsset", "asset113", "yellow", "5", "Tom", "1300")
	if err != nil {
		log.Fatalf("Failed to Submit transaction: %v", err)
	}
	log.Println(string(result))

	log.Println("--> Evaluate Transaction: ReadAsset, function returns an asset with a given assetID")
	/*
		? Read a certain asset
	*/
	result, err = contract.EvaluateTransaction("ReadAsset", "asset113")
	if err != nil {
		log.Fatalf("Failed to evaluate transaction: %v\n", err)
	}
	log.Println(string(result))

	log.Println("--> Evaluate Transaction: AssetExists, function returns 'true' if an asset with given assetID exist")
	/*
		? Check if a certain asset exists
	*/
	result, err = contract.EvaluateTransaction("AssetExists", "asset1")
	if err != nil {
		log.Fatalf("Failed to evaluate transaction: %v\n", err)
	}
	log.Println(string(result))

	log.Println("--> Submit Transaction: TransferAsset asset1, transfer to new owner of Tom")
	/*
		? Transfer the asset
	*/
	_, err = contract.SubmitTransaction("TransferAsset", "asset1", "Tom")
	if err != nil {
		log.Fatalf("Failed to Submit transaction: %v", err)
	}

	log.Println("--> Evaluate Transaction: ReadAsset, function returns 'asset1' attributes")
	/*
		? read another asset
	*/
	result, err = contract.EvaluateTransaction("ReadAsset", "asset1")
	if err != nil {
		log.Fatalf("Failed to evaluate transaction: %v", err)
	}
	log.Println(string(result))
	log.Println("============ application-golang ends ============")
}

// ? Function to populate a certain wallet
func populateWallet(wallet *gateway.Wallet) error {
	log.Println("============ Populating wallet ============")
	//? This construct a path to the msp we desire to inspect
	//? It results in this path : "../../test-network/organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp"
	credPath := filepath.Join(
		"..",
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"users",
		"User1@org1.example.com",
		"msp",
	)
	//? It joins the path until the "cert.pem"
	certPath := filepath.Join(credPath, "signcerts", "cert.pem")
	//? Read the certificate with the cleanest path
	cert, err := os.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}
	//? it joins the path until the "keystore"
	keyDir := filepath.Join(credPath, "keystore")
	// ? we read the dir because there must be atleast 1 private key
	files, err := os.ReadDir(keyDir)
	if err != nil {
		return err
	}
	//? Check the length of the files in it
	if len(files) != 1 {
		return fmt.Errorf("keystore folder should have contain one file")
	}
	//? get the path to the key that is inside of the dir
	keyPath := filepath.Join(keyDir, files[0].Name())
	//? read the key
	key, err := os.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}
	//? Create new identity
	identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))
	//? Put in the wallet the identity "appUser"
	return wallet.Put("appUser", identity)
}
