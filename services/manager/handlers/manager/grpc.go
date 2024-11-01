package handlers

import (
	"context"

	manager "github.com/condemo/raspi-htmx-service/services/common/genproto/services"
	"github.com/condemo/raspi-htmx-service/services/manager/types"
	"google.golang.org/grpc"
)

type ManagerGrpcHandler struct {
	manager.UnimplementedServiceManagerServer
	serviceManager types.ServiceManager
}

func NewManagerGrpcHandler(grpc *grpc.Server, sm types.ServiceManager) {
	gRPCHandler := &ManagerGrpcHandler{serviceManager: sm}

	manager.RegisterServiceManagerServer(grpc, gRPCHandler)
}

func (h *ManagerGrpcHandler) RegisterService(ctx context.Context, req *manager.RegisterServiceRequest) (*manager.RegisterServiceResponse, error) {
	if err := h.serviceManager.RegisterService(ctx, req); err != nil {
		res := &manager.RegisterServiceResponse{
			Message: err.Error(),
		}
		return res, err
	}

	res := &manager.RegisterServiceResponse{
		Message: "success",
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
