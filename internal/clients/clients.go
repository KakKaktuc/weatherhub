package clients

import (
	"encoding/json"
	"fmt"
	"net/http"
	"weatherhub/internal/models"
)

type OpenWeatherClient struct {
	APIKey string
}

func (c *OpenWeatherClient) GetWeather(city string) (models.WeatherData, error) {
	var data models.WeatherData
	url := fmt.Sprintf(
        "https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric",
		city, c.APIKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()

	var result struct {
		Main struct {
			Temp		float64 `json:"temp"`
			Humidity	float64 `json:"humidity"`
		} `json:"main"`
		Wind struct {
			Speed float64 `json:"speed"`
		} `json:"wind"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return data, err
	}

	data = models.WeatherData {
		Source: 		"OpenWeatherMap",
		Temperature:	result.Main.Temp,
		Humidity:		result.Main.Humidity,
		WindSpeed:		result.Wind.Speed,
	}
	return data, nil
}
