package main

import (
	wearablepb "github.com/MarioCarrion/grpc-microservice-example/gen/go/wearable/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type wearableService struct {
	wearablepb.UnimplementedWearableServiceServer
}

// ? This function got changes in order to become a mock
func (w *wearableService) BeatsPerMinute(
	req *wearablepb.BeatsPerMinuteRequest,
	stream wearablepb.WearableService_BeatsPerMinuteServer) error {
	//? Mock object
	results := []uint32{5, 10, 15, 20, 25}

	for i := 0; i < len(results); i++ {
		select {
		//? Cas stream is done
		case <-stream.Context().Done():
			return status.Error(codes.Canceled, "Stream has ended")
		default:
			value := results[i]
			//? Send the messages via stream
			if err := stream.SendMsg(&wearablepb.BeatsPerMinuteResponse{
				Value:  uint32(value + 30),
				Minute: uint32(value),
			}); err != nil {
				return status.Error(codes.Canceled, "Stream has ended")
			}
		}
	}

	return nil
}
