package service

import (
	"context"

	manager "github.com/condemo/raspi-htmx-service/services/common/genproto/services"
)

type ManagerService struct {
	// dependency injection: DB, etc...
	serviceList map[int32]*manager.RaspiService
}

func NewManagerService() *ManagerService {
	return &ManagerService{
		serviceList: make(map[int32]*manager.RaspiService),
	}
}

func (s *ManagerService) LoadService(ctx context.Context, sl *manager.RaspiService) {
	s.serviceList[sl.GetId()] = sl
}

func (s *ManagerService) GetServices(ctx context.Context) []*manager.RaspiService {
	sl := make([]*manager.RaspiService, len(s.serviceList))
	for i, service := range s.serviceList {
		sl[i-1] = service
	}
	return sl
}

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
