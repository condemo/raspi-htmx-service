package service

import (
	"context"
	"fmt"
	"time"

	"github.com/condemo/raspi-htmx-service/services/common/genproto/services/logger"
)

type LoggerService struct {
	// injects..
}

func NewLoggerService() *LoggerService {
	return &LoggerService{}
}

func (s *LoggerService) LogMessage(ctx context.Context, req *logger.LogRequest) error {
	// PERF: Podría user un `strings.Builder` para ser mas eficiente
	var color string

	switch req.GetType() {
	case logger.MessageType_INFO:
		color = COLOR_INFO
	case logger.MessageType_ERROR:
		color = COLOR_ERROR
	case logger.MessageType_SUCCESS:
		color = COLOR_SUCCESS
	case logger.MessageType_WARNING:
		color = COLOR_WARNING
	}

	currentTime := time.Now().Format("01/02/2006 15:04:05")
	fmt.Printf("%s%s [%s%s%s]: %s%s\n", color, currentTime, COLOR_NONE, req.GetServiceName(), color, req.GetMsg(), COLOR_NONE)

	return nil
}
