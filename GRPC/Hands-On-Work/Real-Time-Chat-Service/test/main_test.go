package test

import (
	c "chat/protofiles"
	cImpl "chat/server/liveChatImpl"
	"context"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

var Client c.ChatClient

func TestMain(t *testing.M) {
	//? Create the server
	lis := bufconn.Listen(1920 * 1920)
	defer lis.Close()
	server := grpc.NewServer()
	defer server.Stop()
	c.RegisterChatServer(server, &cImpl.ChatServer{
		Users:       []*c.UserData{},
		Connections: make(map[string]*c.Chat_JoinServer)},
	)
	go func() {
		err := server.Serve(lis)
		if err != nil {
			log.Fatalf("something is not right with the listening %v", err)
		}
	}()
	//? Create the client
	conn, err := grpc.DialContext(
		context.Background(),
		"",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(
			func(ctx context.Context, s string) (net.Conn, error) {
				return lis.Dial()
			},
		),
	)
	if err != nil {
		log.Fatalf("something went wrong with the connection %v", err)
	}
	defer conn.Close()
	Client = c.NewChatClient(conn)
	//? Run the tests
	t.Run()
}
