package types

import (
	"context"

	"github.com/condemo/raspi-htmx-service/services/common/genproto/pb"
)

type RaspiService interface {
	// PERF: Cambiar por completo este implementación haciendola más ajustada a mis necesitades
	// en los handlers se pueden hacer las conversiones necesarias antes de retornar data
	// a otros servicios
	Init(context.Context) error
	Start(context.Context) error
	Stop(context.Context) error
	GetStatus(context.Context) *pb.RaspiService
	GetConfig(context.Context) *pb.ConfigResponse
	UpdateConfig(context.Context, *pb.ServiceConfig) (*pb.ConfigResponse, error)
	GetFullInfo(context.Context, *pb.EmptyRequest) *pb.ServiceFullInfo
}

// PERF: Estructura simple para iniciar, mejorar
type InfoCard struct {
	Icon        string
	Data        string
	LastUpdated string
}
