package inutil

import (
	manager "github.com/condemo/raspi-htmx-service/services/common/genproto/services"
	raspiservices "github.com/condemo/raspi-htmx-service/services/common/genproto/services/raspi_services"
)

func RaspiToManager(in *raspiservices.StatusResponse) *manager.RaspiService {
	return &manager.RaspiService{
		Id: in.GetId(), Name: in.GetName(), Status: in.GetStatus(),
		Data: &manager.ServiceCardData{
			Icon:        in.GetData().GetIcon(),
			DataText:    in.GetData().GetData(),
			LastUpdated: in.GetData().GetLastUpdated(),
		},
	}
}
