package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	u "user-auth/protofiles"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/examples/data"
)

func Register(client u.UserAuthClient, credentials *u.Credentials) error {
	ctx := context.Background()
	confirmation, err := client.Register(ctx, credentials)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	fmt.Println("[Server]", confirmation)
	return nil
}

func Login(client u.UserAuthClient, credentials *u.Credentials) error {
	ctx := context.Background()
	confirmation, err := client.Login(ctx, credentials)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	fmt.Println("[Server]", confirmation)
	return nil
}

func HelloWorld(client u.UserAuthClient, jwt *u.Jwt) error {
	ctx := context.Background()
	confirmation, err := client.HelloWorld(ctx, jwt)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	fmt.Println("[Server]", confirmation)
	return nil
}

func handleLogin(buffer *bufio.Reader, cli u.UserAuthClient) error {
	fmt.Println("*************************************************")
	fmt.Println("Provide the username:")
	fmt.Println("*************************************************")
	username, err := buffer.ReadString('\n')
	if err != nil {
		return fmt.Errorf("error reading the string %v", err)
	}
	username = strings.TrimRight(username, "\n")
	fmt.Println("*************************************************")
	fmt.Println("Provide the password:")
	fmt.Println("*************************************************")
	password, err := buffer.ReadString('\n')
	if err != nil {
		return fmt.Errorf("error reading the string %v", err)
	}
	password = strings.TrimRight(password, "\n")
	credentials := &u.Credentials{}
	credentials.Username = username
	credentials.Password = password
	err = Login(cli, credentials)
	if err != nil {
		return err
	}
	return nil
}

func handleRegister(buffer *bufio.Reader, cli u.UserAuthClient) error {
	fmt.Println("*************************************************")
	fmt.Println("Provide the username:")
	fmt.Println("*************************************************")
	username, err := buffer.ReadString('\n')
	if err != nil {
		return fmt.Errorf("error reading the string %v", err)
	}
	username = strings.TrimRight(username, "\n")
	fmt.Println("*************************************************")
	fmt.Println("Provide the password:")
	fmt.Println("*************************************************")
	password, err := buffer.ReadString('\n')
	if err != nil {
		return fmt.Errorf("error reading the string %v", err)
	}
	password = strings.TrimRight(password, "\n")
	credentials := &u.Credentials{}
	credentials.Username = username
	credentials.Password = password
	err = Register(cli, credentials)
	if err != nil {
		return err
	}
	return nil
}

func handleHelloWorld(buffer *bufio.Reader, cli u.UserAuthClient) error {
	fmt.Println("*************************************************")
	fmt.Println("Provide the jwt:")
	fmt.Println("*************************************************")
	jwt, err := buffer.ReadString('\n')
	if err != nil {
		return fmt.Errorf("error reading the string %v", err)
	}
	jwt = strings.TrimRight(jwt, "\n")
	passport := &u.Jwt{}
	passport.Msg = jwt
	err = HelloWorld(cli, passport)
	if err != nil {
		return err
	}
	return nil
}

// ? Init vars
var (
	tls    = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile = flag.String("ca_file", "", "The file containing the CA root cert file")
)

func main() {
	opts := []grpc.DialOption{}
	flag.Parse()
	if *tls {
		if *caFile == "" {
			*caFile = data.Path("x509/ca_cert.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, "localhost")
		if err != nil {
			log.Fatalf("Failed to create TLS credentials: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	conn, err := grpc.Dial("localhost:2000", opts...)
	if err != nil {
		log.Fatalf("error in dial connection %v", err)
	}
	client := u.NewUserAuthClient(conn)
	fmt.Println("*************************************************")
	fmt.Println("Authentication Service SA")
	fmt.Println("*************************************************")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("*************************************************")
		fmt.Println("1- Login")
		fmt.Println("2- Register")
		fmt.Println("3- HelloWorld Service")
		fmt.Println("4- Exit")
		fmt.Println("*************************************************")
		msg, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error reading the string %v", err)
		}
		msg = strings.TrimRight(msg, "\n")
		if msg == "1" {
			err := handleLogin(reader, client)
			if err != nil {
				log.Fatalf("%v", err)
			}
		} else if msg == "2" {
			err := handleRegister(reader, client)
			if err != nil {
				log.Fatalf("%v", err)
			}
		} else if msg == "3" {
			err := handleHelloWorld(reader, client)
			if err != nil {
				log.Fatalf("%v", err)
			}
		} else if msg == "4" {
			break
		}
	}
}
