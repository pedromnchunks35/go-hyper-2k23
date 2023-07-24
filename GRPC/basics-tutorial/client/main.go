package main

import (
	"context"
	"flag"
	"io"
	"log"
	"math/rand"
	"time"

	pb "basic-tut/proto-files"
)

// ? Init vars
var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

// ? Print feature
func printFeature(client pb.RouteGuideClient, point *pb.Point) {
	log.Printf("Getting feature for point (%d,%d)", point.Latitude, point.Longitude)
	//? Create some context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//? get feature, using the given interface
	feature, err := client.GetFeature(ctx, point)
	if err != nil {
		log.Fatalf("client.GetFeature failed: %v", err)
	}
	log.Println(feature)
}

// ? Print feature's
func printFeatures(client pb.RouteGuideClient, rect *pb.Rectangle) {
	log.Printf("Looking for features withing %v", rect)
	//? Get some context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//? Invoke the method
	stream, err := client.ListFeatures(ctx, rect)
	if err != nil {
		log.Fatalf("client.ListFeatures failed: %v", err)
	}
	//? Get the stream, it comes in a buffer
	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("client.ListFeatures failed: %v", err)
		}
		log.Printf("Feature: name: %q, point: (%v,%v)", feature.GetName(), feature.GetLocation().GetLatitude(), feature.GetLocation().GetLongitude())
	}
}

// ? Sends a sequence of points and expects to receive a summary
func runRecordRoute(client pb.RouteGuideClient) {
	//? Generate random number
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	pointCount := int(r.Int31n(100)) + 2
	var points []*pb.Point
	for i := 0; i < pointCount; i++ {
		points = append(points, randomPoint(r))
	}
	log.Printf("Traversing %d points.", len(points))
	//? create some context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.RecordRoute(ctx)
	if err != nil {
		log.Fatalf("client.RecordRoute failed: %v", err)
	}
	for _, point := range points {
		if err := stream.Send(point); err != nil {
			log.Fatalf("client.RecordRoute: stream.Send(%v) failed: %v", point, err)
		}
	}
	//? Close stream
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("client.RecordRoute failed: %v", err)
	}
	log.Printf("Route summary: %v", reply)
}

//? runRouteChat receives
