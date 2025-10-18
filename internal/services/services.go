package services

import (
	"weatherhub/internal/models"
	"fmt"
)
type WeatherProvider interface {
    GetWeather(city string) (models.WeatherData, error)
}

func AggregateWeather(city string, providers []WeatherProvider) (models.WeatherData, error) {
    var sumTemp, sumHum, sumWind float64
    var count int

    for _, p := range providers {
        data, err := p.GetWeather(city)
        if err != nil {
            continue // можно логировать ошибки
        }
        sumTemp += data.Temperature
        sumHum += data.Humidity
        sumWind += data.WindSpeed
        count++
    }

    if count == 0 {
        return models.WeatherData{}, fmt.Errorf("no data available")
    }

    return models.WeatherData{
        Source:      "Average",
        Temperature: sumTemp / float64(count),
        Humidity:    sumHum / float64(count),
        WindSpeed:   sumWind / float64(count),
    }, nil
}
