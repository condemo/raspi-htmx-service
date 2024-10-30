package service

import (
	"context"

	sysinfo "github.com/condemo/raspi-htmx-service/services/common/genproto/services/sys_info"
	"github.com/condemo/raspi-htmx-service/services/sys_info/tools"
)

type SysInfoService struct {
	// dependency injection
	info *sysinfo.SysInfo
}

func NewSysInfoService() *SysInfoService {
	return &SysInfoService{}
}

// TODO: Implementar
func (s *SysInfoService) GetInfo(ctx context.Context) *sysinfo.SysInfo {
	si := tools.NewSysInfo()
	return si
}
