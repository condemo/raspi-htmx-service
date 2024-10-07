package services

import (
	"errors"
	"fmt"

	"github.com/condemo/raspi-htmx-service/types"
)

// TODO: Otra variable global que podría implementar correctamente
var ServicesList []types.RaspiService

func LoadServices() {
	for i := range 6 {
		s := types.NewRaspiService(fmt.Sprintf("service-%d", i))
		ServicesList = append(ServicesList, s)
	}
}

// TODO: CUTRE
func UpdateService(n string, status bool) (types.RaspiService, error) {
	for i, service := range ServicesList {
		if service.Name == n {
			service.Status = types.ServiceStatus(status)
			ServicesList[i] = service
			return service, nil
		}
	}

	return types.RaspiService{}, errors.New("service not found")
}
