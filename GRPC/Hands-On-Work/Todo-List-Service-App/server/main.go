package main

import (
	"log"
	"net"
	t "tasks/protofiles"
	tmImpl "tasks/server/todoListServerImpl"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:2000")
	if err != nil {
		log.Fatalf("some error occured with the listening of the server %v", err)
	}
	newService := tmImpl.InitTaskManager()
	grpc := grpc.NewServer()
	t.RegisterTaskManagerServer(grpc, newService)
	log.Printf("Server listening at %v", lis.Addr())
	if err := grpc.Serve(lis); err != nil {
		log.Fatalf("error in server creation: %v", err)
	}
}
