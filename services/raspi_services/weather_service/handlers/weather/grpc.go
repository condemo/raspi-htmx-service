package handlers

import (
	"context"
	"log"

	"github.com/condemo/raspi-htmx-service/services/common/config"
	"github.com/condemo/raspi-htmx-service/services/common/genproto/services/logger"
	raspiservices "github.com/condemo/raspi-htmx-service/services/common/genproto/services/raspi_services"
	"github.com/condemo/raspi-htmx-service/services/common/types"
	"github.com/condemo/raspi-htmx-service/services/common/util"
	"google.golang.org/grpc"
)

type WeatherGrpcHandler struct {
	raspiservices.UnimplementedWeatherServiceServer
	wservice   types.RaspiService
	logService logger.LoggerServiceClient
}

func NewWeatherGrpcHandler(grpc *grpc.Server, ws types.RaspiService) {
	logGrpc := util.NewGrpcClient(config.ServicesConfig.LoggerServPort)
	logConn := logger.NewLoggerServiceClient(logGrpc)

	gRPCHandler := &WeatherGrpcHandler{
		wservice:   ws,
		logService: logConn,
	}
	gRPCHandler.wservice.SetLogger(logConn)

	if err := gRPCHandler.wservice.Init(context.Background()); err != nil {
		log.Fatal("error on weather handler init -", err)
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
