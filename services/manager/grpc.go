package manager

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	raspiservices "github.com/condemo/raspi-htmx-service/services/common/genproto/services/raspi_services"
	handlers "github.com/condemo/raspi-htmx-service/services/manager/handlers/manager"
	"github.com/condemo/raspi-htmx-service/services/manager/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	wGrpcClient := newGrpcClient(":8010")
	wConn := raspiservices.NewWeatherServiceClient(wGrpcClient)
	handlers.NewManagerGrpcHandler(gServer, managerService, wConn)

	go func() {
		log.Println("Manager grpc on port", s.addr)
		log.Fatal(gServer.Serve(lis))
	}()

	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-sigC

	gServer.GracefulStop()
}

// PERF: Mover esto a `common` e importar en los servicios que hagan falta
func newGrpcClient(addr string) *grpc.ClientConn {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("error creating grcp client", err)
	}

	return conn
}
