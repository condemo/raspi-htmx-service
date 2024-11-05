package main

import weatherservice "github.com/condemo/raspi-htmx-service/services/raspi_services/weather_service"

func main() {
	grpcServer := weatherservice.NewGrpcServer(":8010")
	grpcServer.Run()
}
