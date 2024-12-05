package types

import (
	"context"

	"github.com/condemo/raspi-htmx-service/services/common/genproto/pb"
)

type Logger interface {
	LogMessage(ctx context.Context, req *pb.LogRequest) error
}
