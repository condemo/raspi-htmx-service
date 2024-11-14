package handlers

import (
	"context"

	"github.com/condemo/raspi-htmx-service/services/common/genproto/services/logger"
	"github.com/condemo/raspi-htmx-service/services/logger/types"
	"google.golang.org/grpc"
)

type LoggerGrpcHandler struct {
	logger.UnimplementedLoggerServiceServer
	loggerService types.Logger
}

func NewLoggerGrpcHandler(grpc *grpc.Server, ls types.Logger) {
	grpcHandler := &LoggerGrpcHandler{
		loggerService: ls,
	}

	logger.RegisterLoggerServiceServer(grpc, grpcHandler)
}

func (h *LoggerGrpcHandler) LogMessage(ctx context.Context, req *logger.LogRequest) (*logger.LogResponse, error) {
	err := h.loggerService.LogMessage(ctx, req)
	if err != nil {
		return nil, err
	}

	res := &logger.LogResponse{}
	return res, nil
}
