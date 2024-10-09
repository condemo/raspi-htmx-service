package types

import (
	"fmt"

	"github.com/uptrace/bun"
)

type ServiceStatus bool

type Service interface {
	Run() error
	Stop() error
}

type RaspiService struct {
	bun.BaseModel `bun:"table:services,alias:sv"`

	Name   string        `bun:"name"`
	Status ServiceStatus `bun:"status,notnull,type:bool"`
	ID     int           `bun:",pk,autoincrement"`
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
	fmt.Println(s.Name, "STOP")
}
