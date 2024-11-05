package service

import (
	"context"
	"fmt"

	raspiservices "github.com/condemo/raspi-htmx-service/services/common/genproto/services/raspi_services"
	"github.com/condemo/raspi-htmx-service/services/raspi_services/weather_service/types"
)

// TODO: Repensar lo que recibe y lo que devueven y volver a implentar
// cambiar la interfaz en `common` en concordancia
type WeatherService struct {
	// injections
	Data types.Weather
}

func NewWeatherService() *WeatherService {
	return &WeatherService{}
}

func (s *WeatherService) Init(ctx context.Context) error {
	s.Data = *types.NewWeather()
	if err := s.Start(ctx); err != nil {
		return err
	}

	return nil
}

func (s *WeatherService) Start(ctx context.Context) error {
	// TODO: Inicia un bucle en una goroutine actualizando la data del tiempo con un ticker

	// ....
	s.Data.State = true
	fmt.Println("Weather Service Working")
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
