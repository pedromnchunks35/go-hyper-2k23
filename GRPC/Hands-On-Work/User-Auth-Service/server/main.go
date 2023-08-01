package main

import (
	"flag"
	"log"
	"net"
	u "user-auth/protofiles"
	uImpl "user-auth/server/authServiceImpl"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/examples/data"
)

var (
	tls      = flag.Bool("tls", false, "Connection use tls if true, else plain TCP")
	certFile = flag.String("cert_file", "", "TLS cert file relative path")
	keyFile  = flag.String("key_file", "", "The TLS key file")
)

func main() {
	flag.Parse()

	bind, err := net.Listen("tcp", "localhost:2000")
	if err != nil {
		log.Fatalf("Something went wrong with the binding %v", err)
	}
	//? Opts
	var opts []grpc.ServerOption
	//? Check if tls is requested as flag
	if *tls {
		//? Check the given attributes
		if *certFile == "" {
			*certFile = data.Path("x509/server_cert.pem")
		}
		if *keyFile == "" {
			*keyFile = data.Path("x509/server_key.pem")
		}
		//? Create credentials from the certificates
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials: %v", err)
		}
		//? add the tls credentials
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpc := grpc.NewServer(opts...)
	u.RegisterUserAuthServer(grpc, &uImpl.UserAuthService{Users: []*u.Credentials{}, JwtSecret: []byte{1, 3, 6, 100, 200, 30, 2, 8}})
	log.Printf("The server will now listen into %v", bind.Addr())
	err = grpc.Serve(bind)
	if err != nil {
		log.Fatalf("Something went wrong serving the server %v", err)
	}
}
