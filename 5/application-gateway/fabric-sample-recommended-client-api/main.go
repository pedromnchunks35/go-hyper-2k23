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

// ? Function to listen events after the last transaction
func eventListener(ctx context.Context, network *client.Network) {
	fmt.Println("\n*** Start chaincode event listening")
	//? Gather events
	events, err := network.ChaincodeEvents(ctx, *chaincodeName)
	if err != nil {
		panic(fmt.Errorf("failed to start chaincode event listening: %w", err))
	}
	//? Print them on the console
	go func() {
		for event := range events {
			asset := formatJSON(event.Payload)
			fmt.Printf("\n<-- Chaincode event received: %s - %s\n", event.EventName, asset)
		}
	}()
}

// ? Function to convert the bytes into json format (in vertical for better readability)
func formatJSON(data []byte) string {
	var result bytes.Buffer
	//? Parse to json in a more readable form
	if err := json.Indent(&result, data, "", "  "); err != nil {
		log.Fatalf("something went wrong parsing the json %v", err)
	}
	return result.String()
}

// ? Function to update a asset with a given contract
func updateAsset(contract *client.Contract) {
	fmt.Printf("\n--> Submit transaction: UpdateAsset, %s update appraised value to 200\n", "1")
	//? Make the update
	_, err := contract.SubmitTransaction("UpdateAsset", "1", "blue", "10", "Sam", "200")
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction: %w", err))
	}

	fmt.Println("\n*** UpdateAsset committed successfully")
}

// ? Creating a asset
func createAsset(contract *client.Contract) uint64 {
	fmt.Printf("\n--> Submit transaction: CreateAsset, %s owned by Sam with appraised value 100\n", "100")
	//? Make a submission that we should await
	_, commit, err := contract.SubmitAsync("CreateAsset", client.WithArguments("100", "blue", "10", "Sam", "100"))
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction: %w", err))
	}
	//? Get the status
	status, err := commit.Status()
	if err != nil {
		panic(fmt.Errorf("failed to get transaction commit status: %w", err))
	}

	if !status.Successful {
		panic(fmt.Errorf("failed to commit transaction with status code %v", status.Code))
	}

	fmt.Println("\n*** CreateAsset committed successfully")
	//? return the block number
	return status.BlockNumber
}

func replayChaincodeEvents(ctx context.Context, network *client.Network, firstBlock uint64) {
	fmt.Println("\n*** Start chaincode event replay***")
	events, err := network.ChaincodeEvents(ctx, *chaincodeName, client.WithStartBlock(firstBlock))
	if err != nil {
		log.Fatalf("error getting events from the first blocl %v", err)
	}
	//? Read the events until a certain phase
	for {
		select {
		case <-time.After(10 * time.Second):
			log.Fatalf("to much time after the await of the event")

		case event := <-events:
			asset := formatJSON(event.Payload)
			fmt.Printf("\n<-- Chaincode event replayed: %s - %s\n", event.EventName, asset)

			if event.EventName == "DeleteAsset" {
				// Reached the last submitted transaction so return to stop listening for events
				return
			}
		}
	}
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
	//? Create a asset that returns the block number for future usage
	firstBlockNumber := createAsset(contract)
	//? Context used for event listening
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//? event listening emitted by subsequent transaction
	eventListener(ctx, network)
	//? update a asset
	updateAsset(contract)
	//? Replay the events from the block containing the first transaction
	replayChaincodeEvents(ctx, network, firstBlockNumber)
}
