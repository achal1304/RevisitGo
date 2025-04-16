package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Weather represents the structure of the weather API response.
type Weather struct {
	Temperature float64 `json:"temperature"`
	Humidity    int     `json:"humidity"`
}

type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

// FetchWeather makes an HTTP call to an external API to fetch the weather data.
func FetchWeather(apiURL string, city string, client HTTPClient) (*Weather, error) {
	// Make the HTTP request to the weather API
	resp, err := client.Get(fmt.Sprintf("%s/weather?city=%s", apiURL, city))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check if the status code is OK
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch weather: %s", resp.Status)
	}

	// Parse the JSON response into the Weather struct
	var weather Weather
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return nil, err
	}

	return &weather, nil
}

// CreateClient initializes an HTTP client
func CreateClient() *http.Client {
	return &http.Client{
		Timeout: 5 * time.Second, // Set a timeout for the HTTP requests
	}
}

func main() {
	apiURL := "http://api.weather.com" // Example API URL (replace with a real one)
	city := "London"

	client := CreateClient()

	weather, err := FetchWeather(apiURL, city, client)
	if err != nil {
		log.Fatalf("Error fetching weather data: %v", err)
	}

	// Print the weather details
	fmt.Printf("Weather in %s: %.2f°C, %d%% humidity\n", city, weather.Temperature, weather.Humidity)
}
