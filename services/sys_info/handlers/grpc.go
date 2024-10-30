package handlers

import (
	"context"

	sysinfo "github.com/condemo/raspi-htmx-service/services/common/genproto/services/sys_info"
	"github.com/condemo/raspi-htmx-service/services/sys_info/types"
	"google.golang.org/grpc"
)

type SysInfoGrpcHandler struct {
	sysinfo.UnimplementedSysInfoServiceServer
	sysInfoService types.SysInfo
}

func NewSysInfoGrpcHandler(grpc *grpc.Server, is types.SysInfo) {
	gRPCHandler := &SysInfoGrpcHandler{sysInfoService: is}

	sysinfo.RegisterSysInfoServiceServer(grpc, gRPCHandler)
}

// TODO:
func (h *SysInfoGrpcHandler) GetInfo(ctx context.Context, req *sysinfo.GetInfoRequest) (*sysinfo.GetInfoResponse, error) {
	si := h.sysInfoService.GetInfo(ctx)
	res := &sysinfo.GetInfoResponse{SisInfo: si}

	return res, nil
}
