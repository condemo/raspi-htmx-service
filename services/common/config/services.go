package config

import "os"

type servicesConf struct {
	ManagerServPort string
	WeatherServPort string
	SysInfoServPort string
	LoggerServPort  string
	WebServPort     string
}

func newServicesConf() *servicesConf {
	return &servicesConf{
		ManagerServPort: os.Getenv("MANAGER_SERVICE_PORT"),
		WeatherServPort: os.Getenv("WEATHER_SERVICE_PORT"),
		SysInfoServPort: os.Getenv("SYSINFO_SERVICE_PORT"),
		LoggerServPort:  os.Getenv("LOGGER_SERVICE_PORT"),
		WebServPort:     os.Getenv("WEB_SERVICE_PORT"),
	}
}

var ServicesConfig = newServicesConf()
