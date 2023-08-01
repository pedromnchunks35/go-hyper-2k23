package test

import (
	"context"
	"testing"
	w "weather/protofiles"
)

func Test_Request_Weather(t *testing.T) {
	req := &w.SearchDetails{}
	req.CountryName = "porto/portugal"
	res, err := Client.RequestWeather(context.Background(), req)
	if err != nil {
		t.Fatalf("error getting server response %v", err)
	}
	if res == nil {
		t.Fatalf("should have a response")
	}
}
