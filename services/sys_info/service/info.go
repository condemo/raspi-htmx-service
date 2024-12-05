package service

import (
	"context"

	"github.com/condemo/raspi-htmx-service/services/common/genproto/pb"
	"github.com/condemo/raspi-htmx-service/services/sys_info/tools"
)

type SysInfoService struct {
	// dependency injection
	info *pb.SysInfo
}

func NewSysInfoService() *SysInfoService {
	return &SysInfoService{}
}

// TODO: Implementar
func (s *SysInfoService) GetInfo(ctx context.Context) *pb.SysInfo {
	si := tools.NewSysInfo()
	return si
}
