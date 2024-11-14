package types

import (
	"context"

	"github.com/condemo/raspi-htmx-service/services/common/genproto/services/logger"
)

type Logger interface {
	LogMessage(ctx context.Context, req *logger.LogRequest) error
}
