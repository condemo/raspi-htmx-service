package main

import "github.com/condemo/raspi-htmx-service/services/manager"

func main() {
	grpcServer := manager.NewGrpcServer(":8000")
	grpcServer.Run()
}
