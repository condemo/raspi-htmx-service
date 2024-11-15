package handlers

import (
	"context"
	"log"

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
	logService logger.LoggerServiceClient
}

func NewWeatherGrpcHandler(grpc *grpc.Server, ws types.RaspiService) {
	logGrpc := util.NewGrpcClient(":7000")
	logConn := logger.NewLoggerServiceClient(logGrpc)

	gRPCHandler := &WeatherGrpcHandler{
		wservice:   ws,
		logService: logConn,
	}
	if err := gRPCHandler.wservice.Init(context.Background()); err != nil {
		log.Fatal("error on weather handler init -", err)
	}

	_, err := gRPCHandler.logService.LogMessage(context.Background(), logs.MakeLog(
		logger.MessageType_SUCCESS, "Weather Handler Starts"))
	if err != nil {
		log.Fatal("error sending log to LogService -", err)
	}

	raspiservices.RegisterWeatherServiceServer(grpc, gRPCHandler)
}

func (h *WeatherGrpcHandler) Start(ctx context.Context, req *raspiservices.EmptyRequest) (*raspiservices.StatusResponse, error) {
	// TODO:
	if err := h.wservice.Start(ctx); err != nil {
		return nil, err
	}

	res := h.wservice.GetStatus(ctx)

	return res, nil
}

func (h *WeatherGrpcHandler) Stop(ctx context.Context, req *raspiservices.EmptyRequest) (*raspiservices.StatusResponse, error) {
	if err := h.wservice.Stop(ctx); err != nil {
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
