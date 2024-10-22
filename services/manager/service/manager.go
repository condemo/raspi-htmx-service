package service

import (
	"context"

	manager "github.com/condemo/raspi-htmx-service/services/common/genproto/services"
)

type ManagerService struct {
	// dependency injection: DB, etc...
}

func NewManagerService() *ManagerService {
	return &ManagerService{}
}

func (s *ManagerService) RegisterService(ctx context.Context, req *manager.RegisterServiceRequest) error {
	// TODO:
	return nil
}

func (s *ManagerService) GetServices(ctx context.Context, req *manager.GetServicesRequest) []*manager.RaspiService {
	// TODO:
	fakeList := []*manager.RaspiService{
		{
			Name: "asadas",
		},
		{
			Name: "kscnskaskdn",
		},
		{
			Name: "caxkcnakska",
		},
	}
	return fakeList
}
