package handlers

import (
	"context"
	"log"

	"github.com/condemo/raspi-htmx-service/services/common/config"
	"github.com/condemo/raspi-htmx-service/services/common/genproto/services/logger"
	raspiservices "github.com/condemo/raspi-htmx-service/services/common/genproto/services/raspi_services"
	"github.com/condemo/raspi-htmx-service/services/common/types"
	"github.com/condemo/raspi-htmx-service/services/common/util"
	"github.com/condemo/raspi-htmx-service/services/raspi_services/weather_service/logs"
	"google.golang.org/grpc"
)

type WeatherGrpcHandler struct {
	raspiservices.UnimplementedWeatherServiceServer
	wservice   types.RaspiService
	LogService logger.LoggerServiceClient
}

func NewWeatherGrpcHandler(grpc *grpc.Server, ws types.RaspiService) *WeatherGrpcHandler {
	logGrpc := util.NewGrpcClient(config.ServicesConfig.LoggerServPort)
	logConn := logger.NewLoggerServiceClient(logGrpc)

	gRPCHandler := &WeatherGrpcHandler{
		wservice:   ws,
		LogService: logConn,
	}

	ctx := context.Background()

	if err := gRPCHandler.wservice.Init(ctx); err != nil {
		log.Fatal("error on weather handler init -", err)
	}

	_, err := gRPCHandler.LogService.LogMessage(ctx, logs.MakeLog(logger.MessageType_SUCCESS, "Service Init"))
	if err != nil {
		log.Fatal("error sending initial weather log")
	}

	raspiservices.RegisterWeatherServiceServer(grpc, gRPCHandler)
	return gRPCHandler
}

func (h *WeatherGrpcHandler) Start(ctx context.Context, req *raspiservices.EmptyRequest) (*raspiservices.StatusResponse, error) {
	// TODO:
	if err := h.wservice.Start(ctx); err != nil {
		return nil, err
	}
	_, err := h.LogService.LogMessage(ctx, logs.MakeLog(logger.MessageType_SUCCESS, "Service ON"))
	if err != nil {
		return nil, err
	}

	res := h.wservice.GetStatus(ctx)

	return res, nil
}

func (h *WeatherGrpcHandler) Stop(ctx context.Context, req *raspiservices.EmptyRequest) (*raspiservices.StatusResponse, error) {
	if err := h.wservice.Stop(ctx); err != nil {
		return nil, err
	}
	_, err := h.LogService.LogMessage(ctx, logs.MakeLog(logger.MessageType_WARNING, "Service OFF"))
	if err != nil {
		return nil, err
	}

	res := h.wservice.GetStatus(ctx)
	return res, nil
}

func (h *WeatherGrpcHandler) GetStatus(ctx context.Context, req *raspiservices.EmptyRequest) (*raspiservices.StatusResponse, error) {
	res := h.wservice.GetStatus(ctx)
	return res, nil
}

func (h *WeatherGrpcHandler) GetFullInfo(ctx context.Context, req *raspiservices.EmptyRequest) (*raspiservices.FullInfoResponse, error) {
	// TODO:
	return nil, nil
}

func (h *WeatherGrpcHandler) GetConfig(ctx context.Context, req *raspiservices.EmptyRequest) (*raspiservices.ConfigResponse, error) {
	// TODO:
	return nil, nil
}

func (h *WeatherGrpcHandler) UpdateConfig(ctx context.Context, req *raspiservices.ConfigRequest) (*raspiservices.ConfigResponse, error) {
	// TODO:
	return nil, nil
}
