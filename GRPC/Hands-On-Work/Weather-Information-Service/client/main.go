package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	w "weather/protofiles"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RequestWeatherInformation(client w.WeatherServiceClient, request *w.SearchDetails) error {
	ctx := context.Background()
	information, err := client.RequestWeather(ctx, request)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	log.Printf("[Server] %v", information)
	return nil
}

var name = flag.String("name", "", "to search a given location for the weather system")

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost:2000", opts...)
	if err != nil {
		log.Fatalf("Error trying to make the connection: %v", err)
	}
	defer conn.Close()
	client := w.NewWeatherServiceClient(conn)
	request := &w.SearchDetails{}
	request.CountryName = *name
	err = RequestWeatherInformation(client, request)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
}
