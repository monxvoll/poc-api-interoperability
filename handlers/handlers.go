package handlers

import (
	"net/http"
	"os"
	"poc-api-interoperability/middleware"
	"poc-api-interoperability/weather"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(r *gin.Engine) {
	// Simple ping route
	r.GET("/ping", PingHandler)
	// Game context route
	r.GET("/game-context", GameContextHandler)
}

// PingHandler handles the /ping route
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// GameContextHandler handles the /game-context route
func GameContextHandler(c *gin.Context) {
	// Try to get city from query parameters, fallback to env variable
	city := c.Query("city")
	if city == "" {
		city = os.Getenv("WEATHER_CITY")
	}
	
	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "City must be provided either as a query parameter '?city=Name' or in the WEATHER_CITY env variable"})
		return
	}

	// 1. Fetch raw data from the external API
	weatherData, err := weather.FetchWeather(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch weather data: " + err.Error()})
		return
	}

	// 2. Translate the raw data into our game domain model
	gameContext := middleware.TranslateWeatherToGameMood(weatherData)

	// 3. Return the standardized game context
	c.JSON(http.StatusOK, gameContext)
}
