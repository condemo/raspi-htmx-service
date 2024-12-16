package service

import (
	"context"
	"fmt"
	"time"

	"github.com/condemo/raspi-htmx-service/services/common/genproto/pb"
	"github.com/condemo/raspi-htmx-service/services/logger/config"
)

type LoggerService struct {
	conf *config.LoggerConfig
}

func NewLoggerService() *LoggerService {
	return &LoggerService{
		conf: config.NewLoggerConfig(),
	}
}

func (s *LoggerService) LogMessage(ctx context.Context, req *pb.LogRequest) error {
	// PERF: Podría usar un `strings.Builder` para ser mas eficiente
	var color string
	currentTime := time.Now().Format("01/02/2006 15:04:05")

	switch req.GetType() {
	case pb.LogMessageType_INFO:
		color = COLOR_INFO
	case pb.LogMessageType_ERROR:
		color = COLOR_ERROR
		msg := fmt.Sprintf("%s [%s]: %s", currentTime, req.GetServiceName(), req.GetMsg())
		if err := s.conf.SaveLog(msg); err != nil {
			return err
		}
	case pb.LogMessageType_SUCCESS:
		color = COLOR_SUCCESS
	case pb.LogMessageType_WARNING:
		color = COLOR_WARNING
	}

	fmt.Printf("%s%s [%s%s%s] %s%s\n", color, currentTime, COLOR_NONE, req.GetServiceName(), color, req.GetMsg(), COLOR_NONE)

	return nil
}

func (s *LoggerService) CleanErrorLog(ctx context.Context, req *pb.CleanErrorReq) error {
	if err := s.conf.Clean(); err != nil {
		return err
	}

	return nil
}
