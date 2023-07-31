package main

import (
	fls "filesys/protofiles"
	flsImp "filesys/server/fileSharingServerImpl"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var PORT = flag.Int("port", 2000, "The port of the server")

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", *PORT))
	if err != nil {
		log.Fatalf("Something went wrong when making the server accessible %v", err)
	}
	grpc := grpc.NewServer()
	fls.RegisterFileSharingServer(grpc, &flsImp.Filesys{})
	log.Printf("Server starting at %v", lis.Addr())
	err = grpc.Serve(lis)
	if err != nil {
		log.Fatalf("Something went wrong making the server available: %v", err)
	}
}
