package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Temperature float64 `json:"temp"`
	} `json:"main"`
	Description []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
}

func getWeather(city string, apiKey string) (*WeatherData, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data WeatherData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func main() {
	apiKey := "3548f44749a6216cacabdf08ba26191d"
	city := "Kraków"

	weather, err := getWeather(city, apiKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Current weather in %s:\n", weather.Name)
	fmt.Printf("Temperature: %.2f°C\n", weather.Main.Temperature)
	fmt.Printf("Conditions: %s - %s\n", weather.Description[0].Main, weather.Description[0].Description)
}
