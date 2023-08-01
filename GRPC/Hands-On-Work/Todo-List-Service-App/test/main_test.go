package test

import (
	"context"
	"log"
	"net"
	t "tasks/protofiles"
	tImpl "tasks/server/todoListServerImpl"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

var Client t.TaskManagerClient

func TestMain(m *testing.M) {
	//? Create server
	lis := bufconn.Listen(1920 * 1920)
	defer lis.Close()
	server := grpc.NewServer()
	defer server.Stop()
	t.RegisterTaskManagerServer(server, tImpl.InitTaskManager())
	go func() {
		err := server.Serve(lis)
		if err != nil {
			log.Fatalf("something went wrong listening the server %v", err)
		}
	}()
	//? Create the client
	conn, err := grpc.DialContext(
		context.Background(),
		"",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
			return lis.Dial()
		}),
	)
	if err != nil {
		log.Fatalf("something went wrong creating the client connection %v", err)
	}
	defer conn.Close()
	Client = t.NewTaskManagerClient(conn)
	//? Run tests
	m.Run()
}
