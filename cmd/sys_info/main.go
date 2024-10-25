package main

import sysinfo "github.com/condemo/raspi-htmx-service/services/sys_info"

func main() {
	grpcServer := sysinfo.NewGrpcServer(":9000")
	grpcServer.Run()
}
