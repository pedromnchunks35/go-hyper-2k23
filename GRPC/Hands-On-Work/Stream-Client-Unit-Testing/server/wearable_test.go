package main

import (
	"context"
	"fmt"
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
	//? Create the server
	server := grpc.NewServer()
	defer server.Stop()
	bind := bufconn.Listen(1920 * 1920)
	defer bind.Close()
	wearablepb.RegisterWearableServiceServer(server, &wearableService{})
	go func() {
		err := server.Serve(bind)
		if err != nil {
			log.Fatalf("something went wrong with the binding %v", err)
		}
	}()
	//? Create the client
	conn, err := grpc.DialContext(
		context.Background(),
		"",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
			return bind.Dial()
		}),
	)
	if err != nil {
		log.Fatalf("something went wrong creating a connection %v", err)
	}
	defer conn.Close()
	client = wearablepb.NewWearableServiceClient(conn)
	t.Run()
}

func Test_BeatsPerMinute(t *testing.T) {
	//? Get the stream
	stream, err := client.ConsumeBeatsPerMinute(context.Background())
	if err != nil {
		t.Fatalf("cannot create the stream %v", err)
	}
	//? Send 5 messages
	for i := 0; i < 5; i++ {
		req := &wearablepb.ConsumeBeatsPerMinuteRequest{}
		req.Minute = uint32(i + 1)
		req.Uuid = fmt.Sprintf("%v", i)
		req.Value = uint32(i * 3)
		err := stream.Send(req)
		if err != nil {
			t.Fatalf("something went wrong with the sending of the request %v", req)
		}
	}
	//? Make the output checks
	total, err := stream.CloseAndRecv()
	if total.Total != uint32(5) {
		t.Fatalf("we only sended 5 numbers")
	}
	if err != nil {
		t.Fatalf("cannot retrieve the message after sending all the stream %v", err)
	}
}
