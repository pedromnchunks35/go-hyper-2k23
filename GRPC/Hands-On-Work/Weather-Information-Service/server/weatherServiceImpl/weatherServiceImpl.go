package weatherServiceImpl

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	w "weather/protofiles"
)

type WeatherService struct {
	*w.UnimplementedWeatherServiceServer
}

func (wsv *WeatherService) RequestWeather(ctx context.Context, request *w.SearchDetails) (*w.Information, error) {
	//? api url
	url := "http://api.weatherapi.com/v1/current.json?"
	apiKey := "key=da9f7a6d164b45bcb5d91031233007"
	searchTerm := fmt.Sprintf("q=%v", request.CountryName)
	req := fmt.Sprintf("%v%v&%v", url, apiKey, searchTerm)
	//? Make the request
	response, err := http.Get(req)
	if err != nil {
		return nil, fmt.Errorf("something went wrong with the request: %v", err)
	}
	if response.StatusCode == 404 {
		return nil, fmt.Errorf("404 not found error")
	}
	defer response.Body.Close()
	//? Get the body
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("something went wrong reading the information of the body %v", err)
	}
	//? Data retrieval
	data := &w.Information{}
	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		return nil, fmt.Errorf("something went wrong unmarshling the data: %v", err)
	}
	//? return the data
	return data, nil
}
