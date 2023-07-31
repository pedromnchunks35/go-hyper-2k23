package main

import (
	"fmt"
	"io"

	wearablepb "github.com/MarioCarrion/grpc-microservice-example/gen/go/wearable/v1"
)

type wearableService struct {
	wearablepb.UnimplementedWearableServiceServer
}

func (w *wearableService) CalculateBeatsPerMinute(stream wearablepb.WearableService_CalculateBeatsPerMinuteServer) error {
	var count, total uint32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		total += req.GetValue()

		fmt.Println("Received", req.GetValue())

		count++

		if count%5 == 0 {
			fmt.Println("Total", total, "Sending", float32(total)/5)

			if err := stream.Send(&wearablepb.CalculateBeatsPerMinuteResponse{
				Average: float32(total) / 5,
			}); err != nil {
				return err
			}

			total = 0
		}
	}
}
