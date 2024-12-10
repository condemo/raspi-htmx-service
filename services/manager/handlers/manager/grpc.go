package handlers

import (
	"context"
	"log"
	"time"

	"github.com/condemo/raspi-htmx-service/services/common/config"
	"github.com/condemo/raspi-htmx-service/services/common/genproto/pb"
	"github.com/condemo/raspi-htmx-service/services/common/util"
	"github.com/condemo/raspi-htmx-service/services/manager/logs"
	"github.com/condemo/raspi-htmx-service/services/manager/types"
	"google.golang.org/grpc"
)

// PERF: Mejorar toda la implementación del handler y del servico/interface, es un lío
type ManagerGrpcHandler struct {
	pb.UnimplementedServiceManagerServer
	serviceManager types.ServiceManager
	logService     pb.LoggerServiceClient
	raspiServices  map[uint32]pb.RaspiServiceClient
}

func NewManagerGrpcHandler(grpc *grpc.Server, sm types.ServiceManager) {
	// Load all the conns
	weatherGrpc := util.NewGrpcClient(config.ServicesConfig.WeatherServPort)
	weatherConn := pb.NewRaspiServiceClient(weatherGrpc)

	logGrpc := util.NewGrpcClient(config.ServicesConfig.LoggerServPort)
	logConn := pb.NewLoggerServiceClient(logGrpc)

	gRPCHandler := &ManagerGrpcHandler{
		serviceManager: sm,
		logService:     logConn,
		raspiServices:  make(map[uint32]pb.RaspiServiceClient),
	}

	//  TODO: podría ser mejor, usar la id que viene del servicio en lugar
	// de hardcodear
	gRPCHandler.raspiServices[0] = weatherConn

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	_, err := gRPCHandler.logService.LogMessage(ctx, logs.MakeLog(
		pb.LogMessageType_SUCCESS, "Manager Handler Starts"))
	if err != nil {
		log.Fatal("error in logger", err)
	}

	// TODO: Load/Read all `RaspiServices` - Cutre
	if err := gRPCHandler.LoadServices(ctx); err != nil {
		_, err := gRPCHandler.logService.LogMessage(ctx, logs.MakeLog(
			pb.LogMessageType_ERROR, "error loading services -"+err.Error()))
		if err != nil {
			log.Fatal("error sending log from manager -", err)
		}
		log.Fatal("error loading services in manager - ", err)
	}

	pb.RegisterServiceManagerServer(grpc, gRPCHandler)
}

// PERF: `LoadServices` no debería estar dentro del handler sino que debaria llamarse
// antes de crear `ManagerGrpcHandler`
func (h *ManagerGrpcHandler) LoadServices(ctx context.Context) error {
	for i := range h.raspiServices {
		ws, err := h.raspiServices[i].GetStatus(ctx, &pb.EmptyRequest{})
		if err != nil {
			h.logService.LogMessage(ctx, logs.MakeLog(
				pb.LogMessageType_ERROR, "error receiving weather data"+err.Error()))
			return err
		}
		h.serviceManager.LoadService(ctx, ws)
	}

	return nil
}

func (h *ManagerGrpcHandler) GetServices(ctx context.Context, req *pb.GetServicesRequest) (*pb.GetServicesResponse, error) {
	err := h.LoadServices(ctx)
	if err != nil {
		return nil, err
	}

	sl := h.serviceManager.GetServices(ctx)
	res := &pb.GetServicesResponse{Services: sl}
	return res, nil
}

func (h *ManagerGrpcHandler) StartService(ctx context.Context, req *pb.ServiceIdRequest) (*pb.RaspiService, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	id := req.GetId()
	st, err := h.raspiServices[id].Start(ctx, &pb.EmptyRequest{})
	if err != nil {
		_, err := h.logService.LogMessage(ctx,
			logs.MakeLog(pb.LogMessageType_ERROR, "error starting weather service -"+err.Error()))
		if err != nil {
			return nil, err
		}
	}
	return st, err
}

func (h *ManagerGrpcHandler) StopService(ctx context.Context, req *pb.ServiceIdRequest) (*pb.RaspiService, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	id := req.GetId()
	st, err := h.raspiServices[id].Stop(ctx, &pb.EmptyRequest{})
	if err != nil {
		_, err := h.logService.LogMessage(ctx,
			logs.MakeLog(pb.LogMessageType_ERROR, "error starting stoping service -"+err.Error()))
		if err != nil {
			return nil, err
		}
	}
	return st, err
}

func (h *ManagerGrpcHandler) GetFullInfo(ctx context.Context, req *pb.ServiceIdRequest) (*pb.ServiceFullInfo, error) {
	id := req.GetId()
	info, err := h.raspiServices[id].GetFullInfo(ctx, &pb.EmptyRequest{})
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (h *ManagerGrpcHandler) GetConfig(ctx context.Context, req *pb.ServiceIdRequest) (*pb.ConfigResponse, error) {
	id := req.GetId()
	res, err := h.raspiServices[id].GetConfig(ctx, &pb.EmptyRequest{})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *ManagerGrpcHandler) UpdateConfig(ctx context.Context, req *pb.ServiceConfig) (*pb.ConfigResponse, error) {
	id := req.GetId()
	res, err := h.raspiServices[id].UpdateConfig(ctx, req)
	return res, err
}
