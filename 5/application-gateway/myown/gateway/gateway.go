package gateway

import (
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

type ConnectionMetadata struct {
	CcpPath string
	Wallet  *gateway.Wallet
	User    string
}

// ? Function to create a connection
func CreateConnection(connectionInfo ConnectionMetadata) (*gateway.Gateway, error) {
	//? Create the connection using the config file and the wallet config
	connection, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(connectionInfo.CcpPath))),
		gateway.WithIdentity(connectionInfo.Wallet, connectionInfo.User),
	)
	if err != nil {
		return nil, err
	}
	return connection, nil
}
