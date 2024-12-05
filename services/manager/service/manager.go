package service

import (
	"context"

	"github.com/condemo/raspi-htmx-service/services/common/genproto/pb"
)

type ManagerService struct {
	// dependency injection: DB, etc...
	serviceList map[int32]*pb.RaspiService
}

func NewManagerService() *ManagerService {
	return &ManagerService{
		serviceList: make(map[int32]*pb.RaspiService),
	}
}

func (s *ManagerService) LoadService(ctx context.Context, sl *pb.RaspiService) {
	s.serviceList[sl.GetId()] = sl
}

func (s *ManagerService) GetServices(ctx context.Context) []*pb.RaspiService {
	sl := make([]*pb.RaspiService, len(s.serviceList))
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
