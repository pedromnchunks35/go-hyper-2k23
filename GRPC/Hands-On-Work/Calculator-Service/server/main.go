package main

import (
	pf "calc/protofiles"
	c "calc/server/calcServerImpl"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var PORT = flag.Int("port", 2000, "the port of the server")

func main() {
	flag.Parse()
	//? Bind to a port
	bind, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", *PORT))
	//? Error handling
	if err != nil {
		log.Fatalf("Failed to do the bind: %v", err)
	}
	//? Create the server
	grpcServer := grpc.NewServer()
	newServer := &c.CalculationServer{}
	pf.RegisterDoMathServer(grpcServer, newServer)
	log.Printf("Server is listening at %v", bind.Addr())
	if err := grpcServer.Serve(bind); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
