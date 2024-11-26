package handlers

import (
	"context"
	"log"

	"github.com/condemo/raspi-htmx-service/services/common/config"
	manager "github.com/condemo/raspi-htmx-service/services/common/genproto/services"
	"github.com/condemo/raspi-htmx-service/services/common/genproto/services/logger"
	raspiservices "github.com/condemo/raspi-htmx-service/services/common/genproto/services/raspi_services"
	"github.com/condemo/raspi-htmx-service/services/common/util"
	"github.com/condemo/raspi-htmx-service/services/manager/inutil"
	"github.com/condemo/raspi-htmx-service/services/manager/logs"
	"github.com/condemo/raspi-htmx-service/services/manager/types"
	"google.golang.org/grpc"
)

// PERF: Mejorar toda la implementación del handler y del servico/interface, es un lío
type ManagerGrpcHandler struct {
	manager.UnimplementedServiceManagerServer
	serviceManager types.ServiceManager
	weatherService raspiservices.WeatherServiceClient
	logService     logger.LoggerServiceClient
}

func NewManagerGrpcHandler(grpc *grpc.Server, sm types.ServiceManager) {
	// Load all the conns
	weatherGrpc := util.NewGrpcClient(config.ServicesConfig.WeatherServPort)
	weatherConn := raspiservices.NewWeatherServiceClient(weatherGrpc)

	logGrpc := util.NewGrpcClient(config.ServicesConfig.LoggerServPort)
	logConn := logger.NewLoggerServiceClient(logGrpc)

	gRPCHandler := &ManagerGrpcHandler{
		serviceManager: sm,
		weatherService: weatherConn,
		logService:     logConn,
	}

	ctx := context.Background()

	_, err := gRPCHandler.logService.LogMessage(ctx, logs.MakeLog(
		logger.MessageType_SUCCESS, "Manager Handler Starts"))
	if err != nil {
		log.Fatal("error in logger", err)
	}

	// TODO: Load/Read all `RaspiServices` - Cutre
	if err := gRPCHandler.LoadServices(context.Background()); err != nil {
		_, err := gRPCHandler.logService.LogMessage(ctx, logs.MakeLog(
			logger.MessageType_ERROR, "error loading services -"+err.Error()))
		if err != nil {
			log.Fatal("error sending log from manager -", err)
		}
		log.Fatal("error loading services in manager - ", err)
	}

	manager.RegisterServiceManagerServer(grpc, gRPCHandler)
}

// PERF: `LoadServices` no debería estar dentro del handler sino que debaria llamarse
// antes de crear `ManagerGrpcHandler`
func (h *ManagerGrpcHandler) LoadServices(ctx context.Context) error {
	ws, err := h.weatherService.GetStatus(ctx, &raspiservices.EmptyRequest{})
	if err != nil {
		h.logService.LogMessage(ctx, logs.MakeLog(
			logger.MessageType_ERROR, "error receiving weather data"+err.Error()))
		return err
	}

	// TODO: Cambiar `LoadService` por `LoadServices`; ir obteniendo la data de cada servicio
	// para luego cargarlo de una con una sola llamada a la función
	h.serviceManager.LoadService(ctx, inutil.RaspiToManager(ws))

	return nil
}

func (h *ManagerGrpcHandler) GetServices(ctx context.Context, req *manager.GetServicesRequest) (*manager.GetServicesResponse, error) {
	err := h.LoadServices(ctx)
	if err != nil {
		return nil, err
	}

	sl := h.serviceManager.GetServices(ctx)
	res := &manager.GetServicesResponse{Services: sl}
	return res, nil
}

func (h *ManagerGrpcHandler) StartService(ctx context.Context, req *manager.ServiceIdRequest) (*manager.RaspiService, error) {
	st, err := h.weatherService.Start(ctx, &raspiservices.EmptyRequest{})
	if err != nil {
		_, err := h.logService.LogMessage(ctx,
			logs.MakeLog(logger.MessageType_ERROR, "error starting weather service -"+err.Error()))
		if err != nil {
			return nil, err
		}
	}
	res := inutil.RaspiToManager(st)
	return res, err
}

func (h *ManagerGrpcHandler) StopService(ctx context.Context, req *manager.ServiceIdRequest) (*manager.RaspiService, error) {
	st, err := h.weatherService.Stop(ctx, &raspiservices.EmptyRequest{})
	if err != nil {
		_, err := h.logService.LogMessage(ctx,
			logs.MakeLog(logger.MessageType_ERROR, "error starting stoping service -"+err.Error()))
		if err != nil {
			return nil, err
		}
	}
	res := inutil.RaspiToManager(st)
	return res, err
}
