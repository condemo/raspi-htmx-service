package types

import (
	"context"

	"github.com/condemo/raspi-htmx-service/services/common/genproto/pb"
)

type SysInfo interface {
	GetInfo(context.Context) *pb.SysInfo
}
