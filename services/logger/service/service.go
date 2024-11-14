package service

import (
	"context"

	"github.com/condemo/raspi-htmx-service/services/common/genproto/services/logger"
)

type LoggerService struct {
	// injects..
}

func NewLoggerService() *LoggerService {
	return &LoggerService{}
}

func (s *LoggerService) LogMessage(ctx context.Context, req *logger.LogRequest) error {
	// TODO:
	return nil
}
