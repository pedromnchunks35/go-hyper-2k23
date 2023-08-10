package main

import (
	"bytes"
	"context"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	tlsCertPath   = flag.String("tls", "", "the tls certificate")
	gatewayPeer   = flag.String("gt_peer_address", "", "the address of the peer")
	certPath      = flag.String("cert", "", "client cert")
	keystorePath  = flag.String("keystore", "", "client keystore")
	mspID         = flag.String("msp_name", "", "the msp we wish to connect to")
	channelName   = flag.String("channel", "", "channel name")
	chaincodeName = flag.String("chaincode", "", "chaincode name")
)

// ? Function to load a certificate
func loadCertificate(filename string) (*x509.Certificate, error) {
	//? Read the given file
	certificatePEM, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file %v", err)
	}
	//? Create a identity from the certificate using the gateway identity package
	return identity.CertificateFromPEM(certificatePEM)
}

// ? Create a new GRPC connection
func newGrpcConnection() *grpc.ClientConn {
	//? load the certificate
	certificate, err := loadCertificate(*tlsCertPath)
	if err != nil {
		log.Fatalf("error loading the certificate %v", certificate)
	}
	//? Creating a certificate pool, which represents the root certificates of CA's
	//? This serves for the client to trust certificates from those sources
	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	//? create credentials for applying it on the grpc client (we need to aapply the peer ip address)
	transportCredentials := credentials.NewClientTLSFromCert(
		certPool,
		*gatewayPeer,
	)
	//? Creathe the connection
	connection, err := grpc.Dial(*gatewayPeer, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		log.Fatalf("errror creating the client connection %v", err)
	}
	return connection
}

// ? Creation of new identity with the respective msp (this identity is for the peer and not for the client connection)
func newIdentity() *identity.X509Identity {
	certificate, err := loadCertificate(*certPath)
	if err != nil {
		log.Fatalf("error loading the certificate %v", err)
	}
	//? Create identity with msp
	id, err := identity.NewX509Identity(
		*mspID,
		certificate,
	)
	if err != nil {
		log.Fatalf("error creating the identity with msp %v", err)
	}
	return id
}

// ? Function to generate a digital signature
func newSign() identity.Sign {
	//? grabbing the directory
	files, err := os.ReadDir(*keystorePath)
	if err != nil {
		log.Fatalf("error reading the keystore dir %v", err)
	}
	//? Getting the privatekeyFile
	privateKeyFile, err := os.ReadFile(path.Join(*keystorePath, files[0].Name()))
	if err != nil {
		log.Fatalf("error reading the private key file %v", err)
	}
	//? Getting the private key itself
	privateKey, err := identity.PrivateKeyFromPEM(privateKeyFile)
	if err != nil {
		log.Fatalf("error getting the private key %v", err)
	}
	//? creating the signature
	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		log.Fatalf("error creating the signature %v", err)
	}
	return sign
}

func eventListener(ctx context.Context, network *client.Network) {
	fmt.Println("\n*** Start chaincode event listening")

	events, err := network.ChaincodeEvents(ctx, *chaincodeName)
	if err != nil {
		panic(fmt.Errorf("failed to start chaincode event listening: %w", err))
	}

	go func() {
		for event := range events {
			asset := formatJSON(event.Payload)
			fmt.Printf("\n<-- Chaincode event received: %s - %s\n", event.EventName, asset)
		}
	}()
}

func formatJSON(data []byte) string {
	var result bytes.Buffer
	if err := json.Indent(&result, data, "", "  "); err != nil {
		panic(fmt.Errorf("failed to parse JSON: %w", err))
	}
	return result.String()
}

func updateAsset(contract *client.Contract) {
	fmt.Printf("\n--> Submit transaction: UpdateAsset, %s update appraised value to 200\n", "1")

	_, err := contract.SubmitTransaction("UpdateAsset", "1", "blue", "10", "Sam", "200")
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction: %w", err))
	}

	fmt.Println("\n*** UpdateAsset committed successfully")
}
func main() {
	//? Create a client connection
	clientConnection := newGrpcConnection()
	defer clientConnection.Close()
	//? Creating a identity
	id := newIdentity()
	sign := newSign()
	//? Creating a sign
	gateway, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		//? Evaluation, endorsement, submittion and commitment timeouts
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		log.Fatalf("error making the connection %v", err)
	}
	defer gateway.Close()
	//? Network and contract get
	network := gateway.GetNetwork(*channelName)
	contract := network.GetContract(*chaincodeName)
	//? Context used for event listening
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//? event listening emitted by subsequent transaction
	eventListener(ctx, network)
	//? update a asset
	updateAsset(contract)
}
