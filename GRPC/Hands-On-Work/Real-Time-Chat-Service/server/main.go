package main

import (
	c "chat/protofiles"
	cImpl "chat/server/liveChatImpl"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:2000")
	if err != nil {
		log.Fatalf("Some error occured listening the server")
	}
	grpc := grpc.NewServer()
	c.RegisterChatServer(
		grpc,
		&cImpl.ChatServer{
			Users:       []*c.UserData{},
			Connections: make(map[string]*c.Chat_JoinServer),
		},
	)
	log.Printf("Server listening at %v", lis.Addr())
	if err := grpc.Serve(lis); err != nil {
		log.Fatalf("Some error occured during the lis of the server %v", err)
	}
}
