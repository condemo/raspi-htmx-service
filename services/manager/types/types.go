package types

import (
	"context"

	manager "github.com/condemo/raspi-htmx-service/services/common/genproto/services"
)

type ServiceManager interface {
	RegisterService(context.Context, *manager.RegisterServiceRequest) error
	GetServices(context.Context, *manager.GetServicesRequest) []*manager.RaspiService
}
