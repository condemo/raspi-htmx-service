package main

import "github.com/condemo/raspi-htmx-service/services/logger"

func main() {
	grpcServer := logger.NewGrpcServer(":7000")
	grpcServer.Run()
}
