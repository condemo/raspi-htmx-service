package logs

import "github.com/condemo/raspi-htmx-service/services/common/genproto/pb"

func MakeLog(prio pb.MessageType, msg string) *pb.LogRequest {
	return &pb.LogRequest{
		ServiceName: pb.ServiceName_WEATHER_SERVICE,
		Type:        prio,
		Msg:         msg,
	}
}
