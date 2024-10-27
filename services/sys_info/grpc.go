package sysinfo

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

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

	go func() {
		log.Println("SysInfo grpc on port", s.addr)
		log.Fatal(gServer.Serve(lis))
	}()

	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-sigC

	gServer.GracefulStop()
}
