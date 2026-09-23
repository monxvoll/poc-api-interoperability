# PoC: API Interoperability & Data Adapter

This Proof of Concept (PoC) is a Go-based service that acts as an interoperability layer. It fetches raw data from an external API and translates it into standardized domain models.

This project explores the Adapter pattern to map varying external data structures into standardized domain models used by a game engine.

## Project Structure

The codebase is organized into distinct packages to separate concerns:

- `main.go`: The application entry point that initializes configuration and the HTTP server.
- `handlers/`: Contains the Gin framework route controllers for the local API.
- `weather/`: The client service responsible for connecting to the OpenWeatherMap API and fetching data.
- `models/`: Defines the core domain models, such as `GameMoodData`, `GameEnvironment`, and `PlayerModifiers`.
- `docs/`: Contains architectural documentation and diagrams .


## Architecture

The system flow relies on a translator (adapter) architecture:
1. The local server receives a request.
2. The `weather` service fetches raw JSON data from OpenWeatherMap.
3. The data is translated into standardized game variables (e.g., Temperature in Celsius mapped to Player Stamina rules).
4. The standardized data is returned to the client to update the game environment.

See the `docs/` directory for detailed sequence and mapping diagrams.


## Prerequisites

- Go 1.20 or higher
- An OpenWeatherMap API key

## Configuration

The application requires environment variables to connect to the external weather API. Create a `.env` file in the root of the project with the following keys:

```env
OPENWEATHER_API_KEY=your_api_key_here
```

## Running the Application

1. Download the required Go modules:
```bash
go mod tidy
```

2. Start the local server:
```bash
go run main.go
```

The server will start on port 8080.

## Endpoints

### 1. Ping
A simple health check endpoint.
```bash
curl http://localhost:8080/ping
```

### 2. Game Context
Fetches the current weather for a specified city and translates it into a game mood context. Requires the `city` query parameter.
```bash
curl "http://localhost:8080/game-context?city=London"
```

**Response Example:**
```json
{
  "city": "London",
  "game_mood": {
    "Theme": "Gloomy",
    "ColorPalette": "Dark_Blue_And_Grey",
    "Environment": {
      "SkyboxTexture": "sky_stormy.png",
      "FogDensity": 0.82,
      "WeatherParticleEffect": "particles_rain"
    },
    "Stats": {
      "StaminaDrainRate": 1,
      "MovementSpeed": 1
    }
  },
  "weather_data": {
    "weather": [...],
    "main": {
      "temp": 14.5,
      "feels_like": 13.2,
      "humidity": 82
    },
    "name": "London"
  }
}
```
