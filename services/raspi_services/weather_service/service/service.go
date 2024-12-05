package service

import (
	"context"
	"sync"
	"time"

	"github.com/condemo/raspi-htmx-service/services/common/genproto/pb"
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
		for {
			select {
			case <-t.C:
				s.mu.RLock()
				s.Data.FullInfo = s.Data.NewFullInfo()
				s.mu.RUnlock()
			case <-s.canChan:
				return
			}
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

func (s *WeatherService) GetStatus(ctx context.Context) *pb.RaspiService {
	ci := s.Data.GetCardInfo()
	return &pb.RaspiService{
		Id: s.id, Status: s.Data.State, Name: s.Data.Name,
		Data: &pb.ServiceCardData{
			Icon:        ci.Icon,
			DataText:    ci.Data,
			LastUpdated: ci.LastUpdated,
		},
	}
}

func (s *WeatherService) GetConfig(ctx context.Context) *pb.ConfigResponse {
	// TODO:
	return nil
}

func (s *WeatherService) UpdateConfig(ctx context.Context, req *pb.ConfigRequest) (*pb.ConfigResponse, error) {
	// TODO:
	return nil, nil
}

func (s *WeatherService) GetFullInfo(ctx context.Context, req *pb.EmptyRequest) *pb.ServiceFullInfo {
	res := &pb.ServiceFullInfo{
		Id: s.id, Name: s.Data.Name, Status: s.Data.State,
		Location: &pb.ServiceLocation{
			City:   s.Data.FullInfo.Location.City,
			Region: s.Data.FullInfo.Location.Region,
			Current: &pb.ServiceCurrentWeather{
				Condition: &pb.ServiceConditionWeather{
					Text: s.Data.FullInfo.Current.Condition.Text,
					Icon: s.Data.FullInfo.Current.Condition.Icon,
				},
				LastUpdated: s.Data.FullInfo.Current.LastUpdated,
				WindDir:     s.Data.FullInfo.Current.WindDir,
				FeelTemp:    s.Data.FullInfo.Current.FeelTemp,
				Temp:        s.Data.FullInfo.Current.Temp,
				WindVel:     s.Data.FullInfo.Current.WindVel,
				IsDay:       s.Data.FullInfo.Current.IsDay,
			},
		},
	}
	return res
}
