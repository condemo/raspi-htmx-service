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

func RaspiFullToManager(in *raspiservices.FullInfoResponse) *manager.ServiceFullInfo {
	return &manager.ServiceFullInfo{
		Id: in.GetId(), Name: in.GetName(), Status: in.GetStatus(),
		Location: &manager.ServiceLocation{
			City:   in.GetLocation().GetCity(),
			Region: in.GetLocation().GetRegion(),
		},
		Current: &manager.ServiceCurrentWeather{
			Condition: &manager.ServiceConditionWeather{
				Text: in.GetCurrent().GetCondition().GetText(),
				Icon: in.GetCurrent().GetCondition().GetIcon(),
			},
			LastUpdated: in.GetCurrent().GetLastUpdated(),
			WindDir:     in.GetCurrent().GetWindDir(),
			WindVel:     in.GetCurrent().GetWindVel(),
			FeelTemp:    in.GetCurrent().GetFeelTemp(),
			Temp:        in.GetCurrent().GetTemp(),
			IsDay:       in.GetCurrent().GetIsDay(),
		},
	}
}
