package services

import (
	"errors"
	"fmt"

	"github.com/condemo/raspi-htmx-service/types"
)

// TODO: Otra variable global que podría implementar correctamente
var ServicesList []types.Service

func LoadServices() {
	ServicesList = append(ServicesList, types.NewWeatherService())
	for i := range 6 {
		s := types.NewDemoService(fmt.Sprintf("Demo-%d", i))
		ServicesList = append(ServicesList, s)
	}
}

// TODO: CUTRE
func UpdateService(n string, status bool) (types.Service, error) {
	for i, service := range ServicesList {
		if service.GetCardData().Name == n {
			service.SwitchStatus(status)
			ServicesList[i] = service
			return service, nil
		}
	}

	return &types.WeatherService{}, errors.New("service not found")
}
