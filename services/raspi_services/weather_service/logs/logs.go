package logs

import "github.com/condemo/raspi-htmx-service/services/common/genproto/services/logger"

func MakeLog(prio logger.MessageType, msg string) *logger.LogRequest {
	return &logger.LogRequest{
		ServiceName: logger.ServiceName_WEATHER_SERVICE,
		Type:        prio,
		Msg:         msg,
	}
}
