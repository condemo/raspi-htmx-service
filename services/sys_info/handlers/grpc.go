package handlers

import (
	"context"
	"log"

	"github.com/condemo/raspi-htmx-service/services/common/config"
	"github.com/condemo/raspi-htmx-service/services/common/genproto/pb"
	"github.com/condemo/raspi-htmx-service/services/common/util"
	"github.com/condemo/raspi-htmx-service/services/sys_info/logs"
	"github.com/condemo/raspi-htmx-service/services/sys_info/types"
	"google.golang.org/grpc"
)

type SysInfoGrpcHandler struct {
	pb.UnimplementedSysInfoServiceServer
	sysInfoService types.SysInfo
	logService     pb.LoggerServiceClient
}

func NewSysInfoGrpcHandler(grpc *grpc.Server, is types.SysInfo) {
	logGrpc := util.NewGrpcClient(config.ServicesConfig.LoggerServPort)
	logConn := pb.NewLoggerServiceClient(logGrpc)

	gRPCHandler := &SysInfoGrpcHandler{
		sysInfoService: is,
		logService:     logConn,
	}

	_, err := gRPCHandler.logService.LogMessage(context.TODO(),
		logs.MakeLog(
			pb.MessageType_SUCCESS, "SysInfo Handler Starts"))
	if err != nil {
		log.Fatal("error sending msg to Logger -", err)
	}

	pb.RegisterSysInfoServiceServer(grpc, gRPCHandler)
}

// TODO:
func (h *SysInfoGrpcHandler) GetInfo(ctx context.Context, req *pb.GetInfoRequest) (*pb.GetInfoResponse, error) {
	si := h.sysInfoService.GetInfo(ctx)
	res := &pb.GetInfoResponse{SisInfo: si}

	return res, nil
}
