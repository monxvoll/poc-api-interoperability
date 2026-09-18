package middleware

import (
	"poc-api-interoperability/models"
	"poc-api-interoperability/weather"
)

// TranslateWeatherToGameMood converts the raw OpenWeatherMap response into a standardized GameMoodData domain model.
func TranslateWeatherToGameMood(w weather.Response) models.GameMoodData {
	var mood models.GameMoodData

	// Get primary weather condition, handle empty slice just in case
	condition := "Clear"
	if len(w.Weather) > 0 {
		condition = w.Weather[0].Main
	}

	// Calculate temperature in Celsius
	tempC := w.Main.Temp - 273.15

	// 1. Map Theme and Color Palette based on condition
	switch condition {
	case "Rain", "Drizzle", "Thunderstorm":
		mood.Theme = "Gloomy"
		mood.ColorPalette = "Dark_Blue_And_Grey"
		mood.Environment.SkyboxTexture = "sky_stormy.png"
		mood.Environment.WeatherParticleEffect = "particles_rain"
	case "Clouds":
		mood.Theme = "Neutral"
		mood.ColorPalette = "Muted_Grey"
		mood.Environment.SkyboxTexture = "sky_cloudy.png"
		mood.Environment.WeatherParticleEffect = "none"
	case "Snow":
		mood.Theme = "Cold"
		mood.ColorPalette = "White_And_Light_Blue"
		mood.Environment.SkyboxTexture = "sky_snow.png"
		mood.Environment.WeatherParticleEffect = "particles_snow"
	case "Clear":
		fallthrough
	default:
		mood.Theme = "Joyful"
		mood.ColorPalette = "Bright_Yellow"
		mood.Environment.SkyboxTexture = "sky_clear.png"
		mood.Environment.WeatherParticleEffect = "none"
	}

	// 2. Map Fog Density based on humidity
	mood.Environment.FogDensity = float32(w.Main.Humidity) / 100.0

	// 3. Map Player Stats based on temperature
	mood.Stats.StaminaDrainRate = 1.0
	mood.Stats.MovementSpeed = 1.0

	if tempC > 30.0 {
		// Hot weather increases stamina drain
		mood.Stats.StaminaDrainRate = 1.5
	} else if tempC < 0.0 {
		// Freezing weather reduces movement speed
		mood.Stats.MovementSpeed = 0.8
	}

	return mood
}
