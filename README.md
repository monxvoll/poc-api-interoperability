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
WEATHER_CITY=London
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
