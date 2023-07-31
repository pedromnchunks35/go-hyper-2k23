package test

import (
	c "calc/protofiles"
	sv "calc/server/calcServerImpl"
	"context"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

var Client c.DoMathClient

func TestMain(t *testing.M) {
	//? Server creation
	server := grpc.NewServer()
	defer server.Stop()
	lis := bufconn.Listen(1920 * 1920)
	c.RegisterDoMathServer(server, &sv.CalculationServer{})
	go func() {
		err := server.Serve(lis)
		if err != nil {
			log.Fatalf("something went wrong with the listening %v", err)
		}
	}()
	//? Create a client
	conn, err := grpc.DialContext(
		context.Background(),
		"",
		grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
			return lis.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("something went wrong making the dial %v", err)
	}
	Client = c.NewDoMathClient(conn)
	//? Run the tests
	t.Run()
}
