package service

import (
	"context"
	"fmt"
	"time"

	"github.com/condemo/raspi-htmx-service/services/common/genproto/pb"
)

type LoggerService struct {
	// injects..
}

func NewLoggerService() *LoggerService {
	return &LoggerService{}
}

func (s *LoggerService) LogMessage(ctx context.Context, req *pb.LogRequest) error {
	// PERF: Podría user un `strings.Builder` para ser mas eficiente
	var color string

	switch req.GetType() {
	case pb.MessageType_INFO:
		color = COLOR_INFO
	case pb.MessageType_ERROR:
		color = COLOR_ERROR
	case pb.MessageType_SUCCESS:
		color = COLOR_SUCCESS
	case pb.MessageType_WARNING:
		color = COLOR_WARNING
	}

	currentTime := time.Now().Format("01/02/2006 15:04:05")
	fmt.Printf("%s%s [%s%s%s] %s%s\n", color, currentTime, COLOR_NONE, req.GetServiceName(), color, req.GetMsg(), COLOR_NONE)

	return nil
}
