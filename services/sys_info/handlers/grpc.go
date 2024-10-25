package handlers

import (
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
