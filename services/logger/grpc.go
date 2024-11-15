package logger

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/condemo/raspi-htmx-service/services/logger/handlers"
	"github.com/condemo/raspi-htmx-service/services/logger/service"
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

	logService := service.NewLoggerService()
	handlers.NewLoggerGrpcHandler(gServer, logService)

	go func() {
		log.Println("Logger grpc on port", s.addr)
		log.Fatal(gServer.Serve(lis))
	}()

	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-sigC

	gServer.GracefulStop()
}
