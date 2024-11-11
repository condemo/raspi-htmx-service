package types

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/condemo/raspi-htmx-service/services/common/types"
	"github.com/condemo/raspi-htmx-service/services/raspi_services/weather_service/config"
)

type Weather struct {
	config   *config.WeatherConfig
	FullInfo *FullInfo
	httpCli  *http.Client
	InfoCard types.InfoCard
	Name     string
	ID       int64
	State    bool
}

func NewWeather() *Weather {
	w := new(Weather)
	w.httpCli = &http.Client{}

	w.Name = "WeatherService"
	w.config = config.NewWeatherConfig()
	w.ID = 1
	w.State = false
	w.FullInfo = w.newFullInfo()

	w.InfoCard = types.InfoCard{
		Icon:        w.FullInfo.Current.Condition.Icon,
		Data:        fmt.Sprintf("%.1f °C", w.FullInfo.Current.Temp),
		LastUpdated: w.FullInfo.Current.LastUpdated,
	}

	return w
}

type FullInfo struct {
	Location struct {
		City   string `json:"name"`
		Region string `json:"region"`
	} `json:"location"`
	Current struct {
		Condition struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
		} `json:"condition"`
		LastUpdated string  `json:"last_updated"`
		WindDir     string  `json:"wind_dir"`
		FeelTemp    float32 `json:"feelslike_c"`
		Temp        float32 `json:"temp_c"`
		WindVel     float32 `json:"wind_kph"`
		IsDay       uint8   `json:"is_day"`
	} `json:"current"`
}

// PERF: Debería devolver un error para poder contestar con un status o error
// al `ManagerService`
func (w *Weather) newFullInfo() *FullInfo {
	fi := new(FullInfo)

	req, err := http.NewRequest(http.MethodGet, "http://api.weatherapi.com/v1/current.json", nil)
	if err != nil {
		log.Fatal("error in newFullInfo ->", err)
	}

	q := req.URL.Query()
	q.Add("key", os.Getenv("WEATHER_API_KEY"))
	q.Add("q", w.config.City)
	q.Add("aqi", "no")
	req.URL.RawQuery = q.Encode()

	resp, err := w.httpCli.Do(req)
	if err != nil {
		log.Fatal("Errored when sending request to the server")
	}

	json.NewDecoder(resp.Body).Decode(fi)
	defer resp.Body.Close()

	return fi
}
