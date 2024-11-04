package weatherservice

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	handlers "github.com/condemo/raspi-htmx-service/services/raspi_services/weather_service/handlers/weather"
	"github.com/condemo/raspi-htmx-service/services/raspi_services/weather_service/service"
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

	weatherservice := service.NewWeatherService()
	handlers.NewWeatherGrpcHandler(gServer, weatherservice)

	go func() {
		log.Println("WeatherService grpc on port", s.addr)
		log.Fatal(gServer.Serve(lis))
	}()

	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-sigC

	gServer.GracefulStop()
}
