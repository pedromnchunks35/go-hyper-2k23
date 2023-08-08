package walletconfig

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

type WalletMetadata struct {
	WalletPath   string
	User         string
	CertPath     string
	KeyStorePath string
	MspName      string
}

// ? Function to populate the wallet
func populateWallet(wallet *gateway.Wallet, walletInfo WalletMetadata) (*gateway.Wallet, error) {
	if walletInfo.CertPath == "" ||
		walletInfo.KeyStorePath == "" ||
		walletInfo.MspName == "" ||
		walletInfo.User == "" {
		return nil, fmt.Errorf("some required fields are missing in order to create a wallet")
	}
	//? Get the cert file
	cert, err := os.ReadFile(filepath.Clean(walletInfo.CertPath))
	if err != nil {
		return nil, err
	}
	//? Read key dir
	keyDir, err := os.ReadDir(walletInfo.KeyStorePath)
	if err != nil {
		return nil, err
	}
	//? Get the key path
	if len(keyDir) != 1 {
		return nil, fmt.Errorf("there is no key inside of the key dir")
	}
	//? Key path
	keyPath := filepath.Join(walletInfo.KeyStorePath, keyDir[0].Name())
	//? Read the key file
	key, err := os.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return nil, err
	}
	//? Create a new identity
	identity := gateway.NewX509Identity(walletInfo.MspName, string(cert), string(key))
	//? put the user inside the identity
	err = wallet.Put(walletInfo.User, identity)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

// ? Function to handle the wallet
func WalletHandler(walletInfo WalletMetadata) (*gateway.Wallet, error) {
	//? Init the wallet
	wallet, err := gateway.NewFileSystemWallet(walletInfo.WalletPath)
	if err != nil {
		return nil, err
	}
	//? Check if the user exists in the wallet
	if !wallet.Exists(walletInfo.User) {
		return populateWallet(wallet, walletInfo)
	}
	//? Simply return the wallet
	return wallet, nil
}
