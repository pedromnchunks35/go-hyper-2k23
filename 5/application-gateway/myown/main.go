package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	g "myown/gateway"
	w "myown/walletconfig"
	"os"
	"strings"

	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

var (
	user         = flag.String("username", "", "the username of the wallet")
	certPath     = flag.String("cert", "", "the path to the cert")
	keystorePath = flag.String("keystore", "", "the path to the keystore")
	mspName      = flag.String("msp_name", "", "the msp name")
	walletPath   = flag.String("wallet_path", "", "the path of the wallet")
	ccpPath      = flag.String("ccp_path", "", "the ccp path")
	test         = flag.Int("test", 0, "")
)

// ? first menu
func mainMenu(reader *bufio.Reader) {
	for {
		fmt.Println("**********************************************************")
		fmt.Println("1- Establish the wallet")
		fmt.Println("exit - Leave the program")
		fmt.Println("**********************************************************")
		option, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("error reading the console %v", err)
		}
		option = strings.TrimRight(option, "\n")
		if option == "exit" {
			break
		} else if option == "1" {
			//? secound menu case the opt is 1
			walletInfo := &w.WalletMetadata{}
			walletInfo.CertPath = *certPath
			walletInfo.KeyStorePath = *keystorePath
			walletInfo.MspName = *mspName
			walletInfo.User = *user
			walletInfo.WalletPath = *walletPath
			wallet, err := w.WalletHandler(*walletInfo)
			if err != nil {
				fmt.Printf("something was wrong with your configuration: %v\n", err)
			} else {
				connectionMenu(reader, wallet)
			}
		}
	}
}

// ? secound menu
func connectionMenu(reader *bufio.Reader, wallet *gateway.Wallet) {
	for {
		fmt.Println("**********************************************************")
		fmt.Println("1- Connect the fabric")
		fmt.Println("exit- Exit this menu")
		fmt.Println("**********************************************************")
		option, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("error reading the console %v", err)
		}
		option = strings.TrimRight(option, "\n")
		if option == "1" {
			connectionInfo := &g.ConnectionMetadata{}
			connectionInfo.CcpPath = *ccpPath
			connectionInfo.User = *user
			connectionInfo.Wallet = wallet
			conn, err := g.CreateConnection(*connectionInfo)
			if err != nil {
				fmt.Printf("something went wrong connecting %v\n", err)
				break
			} else {
				channelMenu(reader, conn)
			}
			defer conn.Close()
		} else if option == "exit" {
			break
		}
	}
}

// ? third menu
func channelMenu(reader *bufio.Reader, conn *gateway.Gateway) {
	for {
		fmt.Println("**********************************************************")
		fmt.Println("1- Connect to a existing channel")
		fmt.Println("exit- Exit this menu")
		fmt.Println("**********************************************************")
		option, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("error reading the console %v", err)
		}
		option = strings.TrimRight(option, "\n")
		if option == "1" {
			fmt.Println("Please provide me the name of the channel you wish to connect")
			option, err := reader.ReadString('\n')
			if err != nil {
				log.Fatalf("error reading the console %v", err)
			}
			option = strings.TrimRight(option, "\n")
			channel, err := conn.GetNetwork(option)
			if err != nil {
				fmt.Printf("something went wrong getting the channel %v \n", err)
				break
			}
			chaincodeMenu(reader, channel)
		} else if option == "exit" {
			break
		}
	}
}

// ? Forth menu
func chaincodeMenu(reader *bufio.Reader, channel *gateway.Network) {
	for {
		fmt.Println("**********************************************************")
		fmt.Println("1- Connect to a existing chaincode")
		fmt.Println("exit- Exit this menu")
		fmt.Println("**********************************************************")
		option, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("error reading the console %v", err)
		}
		option = strings.TrimRight(option, "\n")
		if option == "1" {
			fmt.Println("Please provide me the chaincodeID")
			option, err := reader.ReadString('\n')
			if err != nil {
				log.Fatalf("error reading the console %v", err)
			}
			option = strings.TrimRight(option, "\n")
			contract := channel.GetContract(option)
			if contract == nil {
				fmt.Println("chaincodeid is invalid")
				break
			}
			chaincodeOperationsMenu(reader, contract)
		} else if option == "exit" {
			break
		}
	}
}

// ? Fifth menu
func chaincodeOperationsMenu(reader *bufio.Reader, contract *gateway.Contract) {
	for {
		var res []byte
		fmt.Println("**********************************************************")
		fmt.Println("1- Read operation")
		fmt.Println("2- Write operation")
		fmt.Println("exit- Exit this menu")
		fmt.Println("**********************************************************")
		option, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("error reading the console %v", err)
		}
		option = strings.TrimRight(option, "\n")
		if option == "1" {
			res, err = operation(reader, contract, 1)
		} else if option == "2" {
			res, err = operation(reader, contract, 2)
		} else if option == "exit" {
			break
		}
		fmt.Printf("Response: %v \n", res)
		fmt.Printf("Error: %v \n", err)
	}
}

// ? operation according to the opt
func operation(reader *bufio.Reader, contract *gateway.Contract, op int) ([]byte, error) {
	fmt.Println("provide the method name please")
	method, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("error reading the console %v", err)
	}
	method = strings.TrimRight(method, "\n")
	fmt.Println("provide the arguments with ',', e.g: '1,2,3,4,5'")
	args, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("error reading the console %v", err)
	}
	args = strings.TrimRight(args, "\n")
	if op == 1 {
		return contract.EvaluateTransaction(method, args)
	} else {
		return contract.SubmitTransaction(method, args)
	}
}

func main() {
	flag.Parse()
	fmt.Println("**********************************************************")
	fmt.Println("Welcome to a simple fabric client")
	fmt.Println("**********************************************************")
	if *test == 0 {
		walletMeta := &w.WalletMetadata{}
		walletMeta.CertPath = "./msp/local/certfiles/cert.pem"
		walletMeta.KeyStorePath = "./msp/local/keystore/"
		walletMeta.MspName = "pedro"
		walletMeta.User = "pedro"
		walletMeta.WalletPath = "./localwallet"
		wallet, err := w.WalletHandler(*walletMeta)
		if err != nil {
			log.Fatalf("error setting the wallet")
		}
		connectionInfo := &g.ConnectionMetadata{}
		connectionInfo.CcpPath = "./ccps/ccp-local.json"
		connectionInfo.User = "pedro"
		connectionInfo.Wallet = wallet
		con, err := g.CreateConnection(*connectionInfo)
		if err != nil {
			log.Fatalf("error in the connection %v", err)
		}
		net, err := con.GetNetwork("channel1")
		if err != nil {
			log.Fatalf("error in the getting of the channel1 %v", err)
		}
		contract := net.GetContract("basic_1.0:e4de097efb5be42d96aebc4bde18eea848aad0f5453453ba2aad97f2e41e0d57")
		res, err := contract.EvaluateTransaction("GetAllAssets")
		if err != nil {
			log.Fatalf("error in evaluation %v", err)
		}
		fmt.Println(res)
	} else {
		reader := bufio.NewReader(os.Stdin)
		mainMenu(reader)
	}
}
