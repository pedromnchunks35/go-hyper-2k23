package main

import (
	"log"
	"net"
	w "weather/protofiles"
	wImpl "weather/server/weatherServiceImpl"

	"google.golang.org/grpc"
)

func main() {
	bind, err := net.Listen("tcp", "localhost:2000")
	if err != nil {
		log.Fatalf("Something went wrong when binding the address: %v", err)
	}
	grpc := grpc.NewServer()
	w.RegisterWeatherServiceServer(grpc, &wImpl.WeatherService{})
	log.Printf("Starting the listening at %v", bind.Addr())
	err = grpc.Serve(bind)
	if err != nil {
		log.Fatalf("Something went wrong starting the server: %v", err)
	}
}
