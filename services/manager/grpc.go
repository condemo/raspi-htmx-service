package manager

import (
	"log"
	"net"

	handlers "github.com/condemo/raspi-htmx-service/services/manager/handlers/manager"
	"github.com/condemo/raspi-htmx-service/services/manager/service"
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

	managerService := service.NewManagerService()
	handlers.NewManagerGrpcHandler(gServer, managerService)

	log.Println("HyperMedia grpc on port", s.addr)

	log.Fatal(gServer.Serve(lis))
}
