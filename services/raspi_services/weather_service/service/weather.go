package service

import (
	"context"

	raspiservices "github.com/condemo/raspi-htmx-service/services/common/genproto/services/raspi_services"
)

// TODO: Repensar lo que recibe y lo que devueven y volver a implentar
// cambiar la interfaz en `common` en concordancia
type WeatherService struct{}

func NewWeatherService() *WeatherService {
	return &WeatherService{
		// injections
	}
}

func (s *WeatherService) Init(ctx context.Context) error {
	// TODO:
	return nil
}

func (s *WeatherService) Start(ctx context.Context) error {
	// TODO:
	return nil
}

func (s *WeatherService) Stop(ctx context.Context) error {
	// TODO:
	return nil
}

func (s *WeatherService) GetConfig(ctx context.Context) *raspiservices.ConfigResponse {
	// TODO:
	return nil
}

func (s *WeatherService) UpdateConfig(ctx context.Context, req *raspiservices.ConfigRequest) (*raspiservices.ConfigResponse, error) {
	// TODO:
	return nil, nil
}

func (s *WeatherService) GetCardInfo(ctx context.Context, req *raspiservices.EmptyRequest) *raspiservices.CardInfoResponse {
	// TODO:
	return nil
}

func (s *WeatherService) GetFullInfo(ctx context.Context, req *raspiservices.EmptyRequest) *raspiservices.FullInfoResponse {
	// TODO:
	return nil
}
