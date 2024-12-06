package utils

import "github.com/condemo/raspi-htmx-service/services/common/genproto/pb"

func MakeLog(prio pb.LogMessageType, msg string) *pb.LogRequest {
	return &pb.LogRequest{
		ServiceName: pb.ServiceName_WEB_SERVICE,
		Type:        prio,
		Msg:         msg,
	}
}
