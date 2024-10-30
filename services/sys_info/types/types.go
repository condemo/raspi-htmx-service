package types

import (
	"context"

	sysinfo "github.com/condemo/raspi-htmx-service/services/common/genproto/services/sys_info"
)

type SysInfo interface {
	GetInfo(context.Context) *sysinfo.SysInfo
}
