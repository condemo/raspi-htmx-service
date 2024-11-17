package main

import (
	"github.com/condemo/raspi-htmx-service/services/common/config"
	sysinfo "github.com/condemo/raspi-htmx-service/services/sys_info"
)

func main() {
	grpcServer := sysinfo.NewGrpcServer(config.ServicesConfig.SysInfoServPort)
	grpcServer.Run()
}
