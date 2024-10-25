package sysinfo

import (
	"log"
	"net"

	"github.com/condemo/raspi-htmx-service/services/sys_info/handlers"
	"github.com/condemo/raspi-htmx-service/services/sys_info/service"
	"google.golang.org/grpc"
)

type grpcServer struct {
	addr string
}

func NewGrpcServer(addr string) *grpcServer {
	return &grpcServer{addr: addr}
}

func (s *grpcServer) Run() {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatal(err)
	}

	gServer := grpc.NewServer()
	// Handlers
	sysInfoService := service.NewSysInfoService()
	handlers.NewSysInfoGrpcHandler(gServer, sysInfoService)

	log.Println("SysInfo grpc on port", s.addr)
	log.Fatal(gServer.Serve(lis))
}
