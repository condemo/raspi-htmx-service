package types

import (
	"context"

	"github.com/condemo/raspi-htmx-service/services/common/genproto/pb"
)

type ServiceManager interface {
	LoadService(context.Context, *pb.RaspiService)
	GetServices(context.Context) []*pb.RaspiService
	StartService(context.Context, int32)
	StopService(context.Context, int32)
}
