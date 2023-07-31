package main

import (
	"context"
	"io"
	"log"
	"net"
	"testing"

	wearablepb "github.com/MarioCarrion/grpc-microservice-example/gen/go/wearable/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

var client wearablepb.WearableServiceClient

func TestMain(t *testing.M) {
	//? Instance the server normally
	lis := bufconn.Listen(1024 * 1024)
	defer lis.Close()
	server := grpc.NewServer()
	defer server.Stop()
	wearablepb.RegisterWearableServiceServer(server, &wearableService{})
	go func() {
		err := server.Serve(lis)
		if err != nil {
			log.Fatalf("some error occured making the server listening %v", err)
		}
	}()
	//? Make the client
	conn, err := grpc.DialContext(
		context.Background(),
		"",
		grpc.WithContextDialer(
			func(ctx context.Context, s string) (net.Conn, error) {
				return lis.Dial()
			},
		),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("error making the dial %v", err)
	}
	defer conn.Close()
	client = wearablepb.NewWearableServiceClient(conn)
	t.Run()
}

func Test_Beats(t *testing.T) {
	stream, err := client.BeatsPerMinute(context.Background(), &wearablepb.BeatsPerMinuteRequest{})
	if err != nil {
		t.Fatalf("error calling the method %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("some error occured %v", err)
		}
		if !(res.Minute == 5 || res.Minute == 10 || res.Minute == 15 || res.Minute == 20 || res.Minute == 25) {
			t.Fatalf("the value that comes from the response was not expected")
		}
	}
}
