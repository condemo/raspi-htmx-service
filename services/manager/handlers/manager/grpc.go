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
	weatherService pb.WeatherServiceClient
	logService     pb.LoggerServiceClient
}

func NewManagerGrpcHandler(grpc *grpc.Server, sm types.ServiceManager) {
	// Load all the conns
	weatherGrpc := util.NewGrpcClient(config.ServicesConfig.WeatherServPort)
	weatherConn := pb.NewWeatherServiceClient(weatherGrpc)

	logGrpc := util.NewGrpcClient(config.ServicesConfig.LoggerServPort)
	logConn := pb.NewLoggerServiceClient(logGrpc)

	gRPCHandler := &ManagerGrpcHandler{
		serviceManager: sm,
		weatherService: weatherConn,
		logService:     logConn,
	}

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
	ws, err := h.weatherService.GetStatus(ctx, &pb.EmptyRequest{})
	if err != nil {
		h.logService.LogMessage(ctx, logs.MakeLog(
			pb.LogMessageType_ERROR, "error receiving weather data"+err.Error()))
		return err
	}

	// TODO: Cambiar `LoadService` por `LoadServices`; ir obteniendo la data de cada servicio
	// para luego cargarlo de una con una sola llamada a la función
	h.serviceManager.LoadService(ctx, ws)

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

	st, err := h.weatherService.Start(ctx, &pb.EmptyRequest{})
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

	st, err := h.weatherService.Stop(ctx, &pb.EmptyRequest{})
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
	info, err := h.weatherService.GetFullInfo(ctx, &pb.EmptyRequest{})
	if err != nil {
		return nil, err
	}
	return info, nil
}
