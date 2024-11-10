package types

import (
	"context"

	raspiservices "github.com/condemo/raspi-htmx-service/services/common/genproto/services/raspi_services"
)

type RaspiService interface {
	// PERF: Cambiar por completo este implementación haciendola más ajustada a mis necesitades
	// en los handlers se pueden hacer las conversiones necesarias antes de retornar data
	// a otros servicios
	Init(context.Context) error
	Start(context.Context) error
	Stop(context.Context) error
	GetStatus(context.Context) *raspiservices.StatusResponse
	GetConfig(context.Context) *raspiservices.ConfigResponse
	UpdateConfig(context.Context, *raspiservices.ConfigRequest) (*raspiservices.ConfigResponse, error)
	GetFullInfo(context.Context, *raspiservices.EmptyRequest) *raspiservices.FullInfoResponse
}

// PERF: Estructura simple para iniciar, mejorar
type InfoCard struct {
	Icon        string
	Data        string
	LastUpdated string
}
