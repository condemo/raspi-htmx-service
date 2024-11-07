package types

import (
	"context"

	manager "github.com/condemo/raspi-htmx-service/services/common/genproto/services"
)

type ServiceManager interface {
	LoadService(context.Context, *manager.RaspiService)
	GetServices(context.Context) []*manager.RaspiService
}
