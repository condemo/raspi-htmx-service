package main

import (
	"github.com/condemo/raspi-htmx-service/services/common/config"
	"github.com/condemo/raspi-htmx-service/services/logger"
)

func main() {
	grpcServer := logger.NewGrpcServer(config.ServicesConfig.LoggerServPort)
	grpcServer.Run()
}
