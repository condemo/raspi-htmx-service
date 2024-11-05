package handlers

import (
	"context"
	"log"

	raspiservices "github.com/condemo/raspi-htmx-service/services/common/genproto/services/raspi_services"
	"github.com/condemo/raspi-htmx-service/services/common/types"
	"google.golang.org/grpc"
)

type WeatherGrpcHandler struct {
	raspiservices.UnimplementedWeatherServiceServer
	wservice types.RaspiService
}

func NewWeatherGrpcHandler(grpc *grpc.Server, ws types.RaspiService) {
	gRPCHandler := &WeatherGrpcHandler{wservice: ws}
	_, err := gRPCHandler.Init(context.Background(), &raspiservices.EmptyRequest{})
	if err != nil {
		log.Fatal("error on weather handler init -", err)
	}

	raspiservices.RegisterWeatherServiceServer(grpc, gRPCHandler)
}

func (h *WeatherGrpcHandler) Init(ctx context.Context, req *raspiservices.EmptyRequest) (*raspiservices.StatusResponse, error) {
	if err := h.wservice.Init(ctx); err != nil {
		return nil, err
	}

	res := &raspiservices.StatusResponse{
		Status: "success",
	}
	return res, nil
}

func (h *WeatherGrpcHandler) Start(ctx context.Context, req *raspiservices.EmptyRequest) (*raspiservices.StatusResponse, error) {
	// TODO:
	return nil, nil
}

func (h *WeatherGrpcHandler) Stop(ctx context.Context, req *raspiservices.EmptyRequest) (*raspiservices.StatusResponse, error) {
	// TODO:
	return nil, nil
}

func (h *WeatherGrpcHandler) GetCardInfo(ctx context.Context, req *raspiservices.EmptyRequest) (*raspiservices.CardInfoResponse, error) {
	// TODO:
	return nil, nil
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
