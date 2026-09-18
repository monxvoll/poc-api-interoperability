package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// FetchWeather gets the weather data for a specific city from OpenWeatherMap
func FetchWeather(city string) (Response, error) {
	var data Response
	apiKey := os.Getenv("OPENWEATHER_API_KEY")

	if apiKey == "" {
		return data, fmt.Errorf("missing OPENWEATHER_API_KEY environment variable")
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return data, fmt.Errorf("request error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return data, fmt.Errorf("the server responded with code: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return data, fmt.Errorf("error parsing JSON: %v", err)
	}

	return data, nil
}

// CheckConnection tests the connection to OpenWeatherMap
func CheckConnection() {
	city := os.Getenv("WEATHER_CITY")
	if city == "" {
		fmt.Println("Missing WEATHER_CITY environment variable")
		return
	}

	data, err := FetchWeather(city)
	if err != nil {
		fmt.Printf("connection check failed: %v\n", err)
		return
	}

	celsius := data.Main.Temp - 273.15

	fmt.Printf("city: %s\nweather: %s (%s)\ntemperature (Celsius): %.2f\n",
		data.Name, data.Weather[0].Main, data.Weather[0].Description, celsius)
}
