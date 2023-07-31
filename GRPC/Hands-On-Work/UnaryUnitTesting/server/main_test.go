package main

import (
	"context"
	"log"
	"net"
	"testing"

	userpb "github.com/MarioCarrion/grpc-microservice-example/gen/go/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

// ? Global variable that retains the state of the iteraction
var client userpb.UserServiceClient

// ? This part is related to serve the tests by creating the connections
func TestMain(t *testing.M) {
	//? Creating the SERVER
	lis := bufconn.Listen(1024 * 1024)
	defer lis.Close()
	srv := grpc.NewServer()
	defer srv.Stop()
	svc := userService{}
	//? Registering the service
	userpb.RegisterUserServiceServer(srv, &svc)
	//? Create a go routine to server the server listening
	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("Server error %v", err)
		}
	}()
	//? Creating the Client
	dialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}
	//? Making the dial
	conn, err := grpc.DialContext(
		context.Background(),
		"",
		grpc.WithContextDialer(dialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("error in dial")
	}
	defer conn.Close()
	//? Assigning the client to the global enviroment var
	client = userpb.NewUserServiceClient(conn)
	//? Run the tests
	t.Run()
}

func TestUserService_GetUser(t *testing.T) {
	ctx := context.Background()
	res, err := client.GetUser(ctx, &userpb.GetUserRequest{Uuid: "123"})
	if err != nil {
		t.Fatalf("Some error occurent with the client %v ", err)
	}
	if res.User.Uuid != "123" && res.User.FullName != "Mario" {
		t.Fatalf("the name and uuid are not correct")
	}
}

func TestAnotherFunc(t *testing.T) {
	res, err := client.GetUser(context.Background(), &userpb.GetUserRequest{Uuid: "123"})
	if err != nil {
		t.Fatalf("Some error occurent with the client %v ", err)
	}
	if res.User.Uuid != "123" && res.User.FullName != "Mario" {
		t.Fatalf("the name and uuid are not correct")
	}
}
