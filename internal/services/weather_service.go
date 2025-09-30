package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type WeatherResponse struct {
	Location struct {
		Name    string `json:"name"`
		Region  string `json:"region"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC float64 `json:"temp_c"`
		TempF float64 `json:"temp_f"`
		TempK float64 `json:"temp_k"`
	} `json:"current"`
}

type WeatherService interface {
	GetWeatherByLocation(location string) (float64, float64, float64, error)
}

type weatherService struct {
	client *http.Client
	apiKey string
}

func NewWeatherService() WeatherService {
	return &weatherService{
		client: &http.Client{},
		apiKey: os.Getenv("WEATHER_API_KEY"),
	}
}

func celsiusToFahrenheit(celsius float64) float64 {
	return celsius*1.8 + 32
}

func celsiusToKelvin(celsius float64) float64 {
	return celsius + 273
}

func (s *weatherService) GetWeatherByLocation(location string) (float64, float64, float64, error) {
	if s.apiKey == "" {
		return 0, 0, 0, fmt.Errorf("weather API key not configured")
	}

	encodedLocation := url.QueryEscape(location)
	requestURL := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", s.apiKey, encodedLocation)
	resp, err := s.client.Get(requestURL)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("can not find zipcode")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, 0, 0, fmt.Errorf("can not find zipcode")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("can not find zipcode")
	}

	var weatherResp WeatherResponse
	if err := json.Unmarshal(body, &weatherResp); err != nil {
		return 0, 0, 0, fmt.Errorf("can not find zipcode")
	}

	tempC := weatherResp.Current.TempC

	tempF := celsiusToFahrenheit(tempC)
	tempK := celsiusToKelvin(tempC)

	return tempC, tempF, tempK, nil
}
