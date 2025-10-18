package main

import (
    "fmt"
    "log"
	"os"

	"github.com/joho/godotenv"
    
	"weatherhub/internal/clients"
    "weatherhub/internal/services"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	
    openWeatherKey := os.Getenv("OPENWEATHER_API_KEY")
	if openWeatherKey == "" {
		log.Fatalf("missing OPENWEATHER_API_KEY")
	}

	openWeather := &clients.OpenWeatherClient{APIKey: openWeatherKey}
    // TODO: добавить других провайдеров

    city := "Otradnoye"

    result, err := services.AggregateWeather(city, []services.WeatherProvider{openWeather})
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Average weather in %s:\nTemp: %.1f°C\nHumidity: %.1f%%\nWind: %.1f m/s\n",
        city, result.Temperature, result.Humidity, result.WindSpeed)
}
