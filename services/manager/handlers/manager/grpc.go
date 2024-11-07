package handlers

import (
	"context"
	"log"

	manager "github.com/condemo/raspi-htmx-service/services/common/genproto/services"
	raspiservices "github.com/condemo/raspi-htmx-service/services/common/genproto/services/raspi_services"
	"github.com/condemo/raspi-htmx-service/services/manager/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ManagerGrpcHandler struct {
	manager.UnimplementedServiceManagerServer
	serviceManager types.ServiceManager
	weatherService raspiservices.WeatherServiceClient
}

func NewManagerGrpcHandler(grpc *grpc.Server, sm types.ServiceManager) {
	// Load all the conns
	weatherGrpc := newGrpcClient(":8010")
	weatherConn := raspiservices.NewWeatherServiceClient(weatherGrpc)

	gRPCHandler := &ManagerGrpcHandler{
		serviceManager: sm,
		weatherService: weatherConn,
	}

	// TODO: Load/Read all `RaspiServices`
	if err := gRPCHandler.LoadServices(context.Background()); err != nil {
		log.Fatal("error loading services in manager - ", err)
	}

	manager.RegisterServiceManagerServer(grpc, gRPCHandler)
}

// PERF: `LoadServices` no debería estar dentro del handler sino que debaria llamarse
// antes de crear `ManagerGrpcHandler`
func (h *ManagerGrpcHandler) LoadServices(ctx context.Context) error {
	ws, err := h.weatherService.GetStatus(ctx, &raspiservices.EmptyRequest{})
	if err != nil {
		return err
	}

	// TODO: Cambiar `LoadService` por `LoadServices`; ir obteniendo la data de cada servicio
	// para luego cargarlo de una con una sola llamada a la función
	h.serviceManager.LoadService(ctx, &manager.RaspiService{Id: ws.Id, Status: ws.Status, Name: ws.Name})

	// TODO: BORRAR
	h.serviceManager.GetServices(ctx)

	return nil
}

func (h *ManagerGrpcHandler) GetServices(ctx context.Context, req *manager.GetServicesRequest) (*manager.GetServicesResponse, error) {
	sl := h.serviceManager.GetServices(ctx)
	res := &manager.GetServicesResponse{Services: sl}
	return res, nil
}

// PERF: Mover esto a `common` e importar en los servicios que hagan falta
func newGrpcClient(addr string) *grpc.ClientConn {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("error creating grcp client", err)
	}

	return conn
}
