package main

import (
	"github.com/condemo/raspi-htmx-service/services/common/config"
	"github.com/condemo/raspi-htmx-service/services/manager"
)

func main() {
	grpcServer := manager.NewGrpcServer(config.ServicesConfig.ManagerServPort)
	grpcServer.Run()
}
