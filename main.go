package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func requestWeatherApi(city string) string {
	api_key := os.Getenv("WEATHER_API_KEY")
	url := "https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + api_key

	// HTTP Get Request
	resp, err := http.Get(url)
	// Error Handling
	if err != nil {
		fmt.Println("Request Error:")
		print(err)
	}
	// Close Request
	defer resp.Body.Close()

	// Read Request Response
	body, err := io.ReadAll(resp.Body)
	// Error Handling
	if err != nil {
		fmt.Println("Read Response Error:")
		print(err)
	}
	// Response
	return string(body)
}

func main() {
	// UI Script
	fmt.Println("---------------------------------------")
	fmt.Println("City Name:")
	// Read Input User
	var city string
	fmt.Scanln(&city)
	fmt.Println("---------------------------------------")
	// Response
	fmt.Println("The Weather in that city is:")
	response := requestWeatherApi(city)
	fmt.Println(response)
}
