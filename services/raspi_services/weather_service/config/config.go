package config

import (
	"github.com/joho/godotenv"
)

type WeatherConfig struct {
	City string
}

func NewWeatherConfig() *WeatherConfig {
	godotenv.Load("services/raspi_services/weather_service/.env")
	return &WeatherConfig{
		City: "zaragoza",
	}
}
