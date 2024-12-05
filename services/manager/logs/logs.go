package logs

import (
	"github.com/condemo/raspi-htmx-service/services/common/genproto/pb"
)

func MakeLog(prio pb.LogMessageType, msg string) *pb.LogRequest {
	return &pb.LogRequest{
		ServiceName: pb.ServiceName_SERVICE_MANAGER,
		Type:        prio,
		Msg:         msg,
	}
}
