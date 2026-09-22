package handlers

import (
	"net/http"
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
	// Get city from query parameters
	city := c.Query("city")
	
	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "City must be provided as a query parameter '?city=Name'"})
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

	// 3. Create a combined response
	response := gin.H{
		"city":         weatherData.Name,
		"weather_data": weatherData,
		"game_mood":    gameContext,
	}

	// 4. Return the standardized game context and actual weather
	c.JSON(http.StatusOK, response)
}
