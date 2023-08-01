package test

import (
	"context"
	"log"
	"net"
	"testing"
	u "user-auth/protofiles"
	uImpl "user-auth/server/authServiceImpl"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

var Client u.UserAuthClient

func TestMain(t *testing.M) {
	//? Create the server
	lis := bufconn.Listen(1920 * 1920)
	defer lis.Close()
	server := grpc.NewServer()
	defer server.Stop()
	u.RegisterUserAuthServer(server, &uImpl.UserAuthService{Users: []*u.Credentials{}, JwtSecret: []byte{1, 3, 6, 100, 200, 30, 2, 8}})
	go func() {
		err := server.Serve(lis)
		if err != nil {
			log.Fatalf("something went wrong listening %v", err)
		}
	}()
	//? Create a client
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
	Client = u.NewUserAuthClient(conn)
	//? Run tests
	t.Run()
}
