package main

import (
	"github.com/condemo/raspi-htmx-service/services/common/config"
	weatherservice "github.com/condemo/raspi-htmx-service/services/raspi_services/weather_service"
)

func main() {
	grpcServer := weatherservice.NewGrpcServer(config.ServicesConfig.WeatherServPort)
	grpcServer.Run()
}
