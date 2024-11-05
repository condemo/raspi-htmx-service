package types

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/condemo/raspi-htmx-service/services/raspi_services/weather_service/config"
)

type Weather struct {
	config   *config.WeatherConfig
	FullInfo *FullInfo
	Name     string
	ID       int64
	State    bool
}

func NewWeather() *Weather {
	w := new(Weather)

	w.Name = "WeatherService"
	w.config = config.NewWeatherConfig()
	w.ID = 1
	w.State = false
	w.FullInfo = newFullInfo(w.config.City)

	return w
}

type CardInfo struct{}

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

func newFullInfo(city string) *FullInfo {
	// PERF: Debería devolver un error para poder contestar con un status o error
	// al `ManagerService`
	fi := new(FullInfo)
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, "http://api.weatherapi.com/v1/current.json", nil)
	if err != nil {
		log.Fatal("error in newFullInfo ->", err)
	}

	q := req.URL.Query()
	q.Add("key", os.Getenv("WEATHER_API_KEY"))
	q.Add("q", city)
	q.Add("aqi", "no")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Errored when sending request to the server")
	}

	json.NewDecoder(resp.Body).Decode(fi)
	defer resp.Body.Close()

	return fi
}
