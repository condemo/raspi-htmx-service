package handlers

import (
	"context"
	"log"

	"github.com/condemo/raspi-htmx-service/services/common/genproto/services/logger"
	sysinfo "github.com/condemo/raspi-htmx-service/services/common/genproto/services/sys_info"
	"github.com/condemo/raspi-htmx-service/services/common/util"
	"github.com/condemo/raspi-htmx-service/services/sys_info/logs"
	"github.com/condemo/raspi-htmx-service/services/sys_info/types"
	"google.golang.org/grpc"
)

type SysInfoGrpcHandler struct {
	sysinfo.UnimplementedSysInfoServiceServer
	sysInfoService types.SysInfo
	logService     logger.LoggerServiceClient
}

func NewSysInfoGrpcHandler(grpc *grpc.Server, is types.SysInfo) {
	logGrpc := util.NewGrpcClient(":7000")
	logConn := logger.NewLoggerServiceClient(logGrpc)

	gRPCHandler := &SysInfoGrpcHandler{
		sysInfoService: is,
		logService:     logConn,
	}

	_, err := gRPCHandler.logService.LogMessage(context.TODO(),
		logs.MakeLog(
			logger.MessageType_SUCCESS, "SysInfo Handler Starts"))
	if err != nil {
		log.Fatal("error sending msg to Logger -", err)
	}

	sysinfo.RegisterSysInfoServiceServer(grpc, gRPCHandler)
}

// TODO:
func (h *SysInfoGrpcHandler) GetInfo(ctx context.Context, req *sysinfo.GetInfoRequest) (*sysinfo.GetInfoResponse, error) {
	si := h.sysInfoService.GetInfo(ctx)
	res := &sysinfo.GetInfoResponse{SisInfo: si}

	return res, nil
}
