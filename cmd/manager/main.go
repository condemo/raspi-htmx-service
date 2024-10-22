package main

import "github.com/condemo/raspi-htmx-service/services/manager"

func main() {
	httpServer := manager.NewHttpServer(":8000")
	go httpServer.Run()

	grpcServer := manager.NewGrpcServer(":8080")
	grpcServer.Run()
}
