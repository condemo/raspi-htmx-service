package service

import (
	"context"
	"sync"
	"time"

	raspiservices "github.com/condemo/raspi-htmx-service/services/common/genproto/services/raspi_services"
	"github.com/condemo/raspi-htmx-service/services/raspi_services/weather_service/types"
)

// TODO: Repensar lo que recibe y lo que devueven y volver a implentar
// cambiar la interfaz en `common` en concordancia
type WeatherService struct {
	mu      *sync.RWMutex
	canChan chan struct{}
	Data    types.Weather
	id      int32
}

func NewWeatherService() *WeatherService {
	return &WeatherService{
		id:      1,
		mu:      new(sync.RWMutex),
		canChan: make(chan struct{}),
	}
}

func (s *WeatherService) Init(ctx context.Context) error {
	s.Data = *types.NewWeather()
	if err := s.Start(ctx); err != nil {
		return err
	}
	return nil
}

func (s *WeatherService) Start(ctx context.Context) error {
	// TODO: Mover la duración a la config para poder modificarla
	t := time.NewTicker(time.Minute * 5)

	go func() {
		select {
		case <-t.C:
			s.mu.RLock()
			s.Data.FullInfo = s.Data.NewFullInfo()
			s.mu.RUnlock()
		case <-s.canChan:
			break
		}
	}()

	s.Data.State = true
	return nil
}

func (s *WeatherService) Stop(ctx context.Context) error {
	s.canChan <- struct{}{}
	s.Data.State = false
	return nil
}

func (s *WeatherService) GetStatus(ctx context.Context) *raspiservices.StatusResponse {
	return &raspiservices.StatusResponse{
		Id: s.id, Status: s.Data.State, Name: s.Data.Name,
		Data: &raspiservices.WeatherCardData{
			Icon:        s.Data.GetCardInfo().Icon,
			Data:        s.Data.GetCardInfo().Data,
			LastUpdated: s.Data.GetCardInfo().LastUpdated,
		},
	}
}

func (s *WeatherService) GetConfig(ctx context.Context) *raspiservices.ConfigResponse {
	// TODO:
	return nil
}

func (s *WeatherService) UpdateConfig(ctx context.Context, req *raspiservices.ConfigRequest) (*raspiservices.ConfigResponse, error) {
	// TODO:
	return nil, nil
}

func (s *WeatherService) GetFullInfo(ctx context.Context, req *raspiservices.EmptyRequest) *raspiservices.FullInfoResponse {
	// TODO:
	return nil
}
