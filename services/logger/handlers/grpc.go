package handlers

import (
	"context"

	"github.com/condemo/raspi-htmx-service/services/common/genproto/pb"
	"github.com/condemo/raspi-htmx-service/services/logger/types"
	"google.golang.org/grpc"
)

type LoggerGrpcHandler struct {
	pb.UnimplementedLoggerServiceServer
	loggerService types.Logger
}

func NewLoggerGrpcHandler(grpc *grpc.Server, ls types.Logger) {
	grpcHandler := &LoggerGrpcHandler{
		loggerService: ls,
	}

	pb.RegisterLoggerServiceServer(grpc, grpcHandler)
}

func (h *LoggerGrpcHandler) LogMessage(ctx context.Context, req *pb.LogRequest) (*pb.LogResponse, error) {
	err := h.loggerService.LogMessage(ctx, req)
	if err != nil {
		return nil, err
	}

	res := &pb.LogResponse{}
	return res, nil
}
