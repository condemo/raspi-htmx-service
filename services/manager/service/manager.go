package service

import (
	"context"

	manager "github.com/condemo/raspi-htmx-service/services/common/genproto/services"
)

type ManagerService struct {
	// dependency injection: DB, etc...
	serviceList []*manager.RaspiService
}

func NewManagerService() *ManagerService {
	return &ManagerService{}
}

func (s *ManagerService) LoadService(ctx context.Context, sl *manager.RaspiService) {
	s.serviceList = append(s.serviceList, sl)
}

func (s *ManagerService) GetServices(ctx context.Context) []*manager.RaspiService {
	return s.serviceList
}

// TODO: Todas las funcionalidades
func (s *ManagerService) StartService(ctx context.Context, id int32) {
	for _, s := range s.serviceList {
		if s.Id == id {
			s.Status = true
		}
	}
}

func (s *ManagerService) StopService(ctx context.Context, id int32) {
	for _, s := range s.serviceList {
		if s.Id == id {
			s.Status = false
		}
	}
}
