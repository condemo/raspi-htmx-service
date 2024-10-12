package types

import (
	"fmt"
)

type ServiceStatus bool

type Service interface {
	Run() error
	Stop() error
	GetFullData() ServiceFullData
	GetCardData() ServiceCardData
	SwitchStatus(bool)
}

type (
	ServiceFullData struct{}
	ServiceCardData struct {
		Name   string
		Status ServiceStatus
	}
)

type WeatherService struct {
	Name   string
	Status ServiceStatus
	ID     int64
}

func NewWeatherService() *WeatherService {
	return &WeatherService{
		Name: "WeatherService",
	}
}

func (s WeatherService) Run() error {
	// TODO:
	fmt.Println(s.Name, "RUN")
	return nil
}

func (s *WeatherService) SwitchStatus(st bool) {
	s.Status = ServiceStatus(st)
}

func (s WeatherService) GetCardData() ServiceCardData {
	return ServiceCardData{
		Name:   s.Name,
		Status: s.Status,
	}
}

func (s WeatherService) GetFullData() ServiceFullData {
	return ServiceFullData{}
}

func (s WeatherService) Stop() error {
	// TODO:
	fmt.Println(s.Name, "STOP")
	return nil
}

// Demo
type DemoService struct {
	Name   string
	Status ServiceStatus
	ID     int64
}

func NewDemoService(name string) *DemoService {
	return &DemoService{
		Name: name,
	}
}

func (s DemoService) Run() error {
	// TODO:
	return nil
}

func (s DemoService) Stop() error {
	// TODO:
	return nil
}

func (s DemoService) GetFullData() ServiceFullData {
	return ServiceFullData{}
}

func (s DemoService) GetCardData() ServiceCardData {
	return ServiceCardData{
		Name:   s.Name,
		Status: s.Status,
	}
}

func (s *DemoService) SwitchStatus(status bool) {
	s.Status = ServiceStatus(status)
}
