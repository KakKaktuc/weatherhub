package models

type WeatherData struct {
	Source string
	Temperature float64 // в C
	Humidity float64 // %
	WindSpeed float64 // м/с
}