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
	//? create server
	server := grpc.NewServer()
	defer server.Stop()
	wearablepb.RegisterWearableServiceServer(server, &wearableService{})
	lis := bufconn.Listen(1920 * 1920)
	defer lis.Close()
	go func() {
		err := server.Serve(lis)
		if err != nil {
			log.Fatalf("something went wrong listening the new server %v", err)
		}
	}()
	//? create client
	conn, err := grpc.DialContext(
		context.Background(),
		"",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
			return lis.Dial()
		}),
	)
	if err != nil {
		log.Fatalf("something went wrong in the dial %v", err)
	}
	client = wearablepb.NewWearableServiceClient(conn)
	//? start the testing
	t.Run()
}

func Test_CalculateTheBeats(t *testing.T) {
	stream, err := client.CalculateBeatsPerMinute(context.Background())
	if err != nil {
		t.Fatalf("something went wrong with the stream %v", err)
	}
	total := 0
	for i := 1; i < 11; i++ {
		request := &wearablepb.CalculateBeatsPerMinuteRequest{}
		request.Minute = uint32(i * 60)
		request.Uuid = fmt.Sprintf("%v", i)
		request.Value = uint32(i + 2)
		err := stream.Send(request)
		if err != nil {
			t.Fatalf("something went wrong sending the request %v", err)
		}
		total += int(request.Value)
		if i != 0 && i%5 == 0 {
			fmt.Println(total)
			res, err := stream.Recv()
			if err != nil {
				t.Fatalf("something went wrong receiving the request %v", err)
			}
			if float32(total)/5 != res.Average {
				t.Fatalf("expected %v, we have %v", float32(total)/5, res.Average)
			}
			total = 0
		}
	}
}
