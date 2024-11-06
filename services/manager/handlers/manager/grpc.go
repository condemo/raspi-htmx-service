package handlers

import (
	"context"
	"log"

	manager "github.com/condemo/raspi-htmx-service/services/common/genproto/services"
	raspiservices "github.com/condemo/raspi-htmx-service/services/common/genproto/services/raspi_services"
	"github.com/condemo/raspi-htmx-service/services/manager/types"
	"google.golang.org/grpc"
)

type ManagerGrpcHandler struct {
	manager.UnimplementedServiceManagerServer
	serviceManager types.ServiceManager
	weatherConn    raspiservices.WeatherServiceClient
}

func NewManagerGrpcHandler(grpc *grpc.Server, sm types.ServiceManager, wConn raspiservices.WeatherServiceClient) {
	gRPCHandler := &ManagerGrpcHandler{
		serviceManager: sm,
		weatherConn:    wConn,
	}

	manager.RegisterServiceManagerServer(grpc, gRPCHandler)
}

func (h *ManagerGrpcHandler) RegisterService(ctx context.Context, req *manager.RegisterServiceRequest) (*manager.ServiceStatusResponse, error) {
	// TODO: Cambiar al reestructurar como inician los servicios
	st, err := h.weatherConn.Start(ctx, &raspiservices.EmptyRequest{})
	if err != nil {
		log.Fatal("error init weather service", err)
	}

	res := &manager.ServiceStatusResponse{
		Message: st.Status,
	}

	return res, nil
}

func (h *ManagerGrpcHandler) GetServices(ctx context.Context, req *manager.GetServicesRequest) (*manager.GetServicesResponse, error) {
	sl := h.serviceManager.GetServices(ctx, req)

	res := &manager.GetServicesResponse{
		Services: sl,
	}
	return res, nil
}

func (h *ManagerGrpcHandler) GetServiceData(ctx context.Context, req *manager.ServiceIdRequest) (*manager.RaspiService, error) {
	return nil, nil
}

func (h *ManagerGrpcHandler) StartService(ctx context.Context, req *manager.ServiceIdRequest) (*manager.ServiceStatusResponse, error) {
	return nil, nil
}

func (h *ManagerGrpcHandler) StopService(ctx context.Context, req *manager.ServiceIdRequest) (*manager.ServiceStatusResponse, error) {
	return nil, nil
}
