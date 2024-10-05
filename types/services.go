package types

import "fmt"

type serviceStatus bool

type Service interface {
	Run() error
	Stop() error
}

type RaspiService struct {
	Name   string
	Status serviceStatus
}

func NewRaspiService(name string) RaspiService {
	return RaspiService{
		Name: name,
	}
}

func (s *RaspiService) Run() {
	// TODO:
	fmt.Println(s.Name, "RUN")
}

func (s *RaspiService) Stop() {
	// TODO:
	fmt.Println(s.Name, "Stop")
}
