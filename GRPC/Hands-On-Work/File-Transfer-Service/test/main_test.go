package test

import (
	"context"
	fls "filesys/protofiles"
	flsImpl "filesys/server/fileSharingServerImpl"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

var Client fls.FileSharingClient

func TestMain(t *testing.M) {
	//? create server
	server := grpc.NewServer()
	defer server.Stop()
	lis := bufconn.Listen(1920 * 1920)
	defer lis.Close()
	fls.RegisterFileSharingServer(server, &flsImpl.Filesys{})
	go func() {
		err := server.Serve(lis)
		if err != nil {
			log.Fatalf("server fake lis error %v", err)
		}
	}()
	//? Create a client
	conn, err := grpc.DialContext(
		context.Background(),
		"",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
			return lis.Dial()
		}),
	)
	if err != nil {
		log.Fatalf("error making the dial %v", err)
	}
	Client = fls.NewFileSharingClient(conn)
	//? Run tests
	t.Run()
}
