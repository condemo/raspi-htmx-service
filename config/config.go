package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var Envs = initEnvConfig()

type envConfig struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

func initEnvConfig() envConfig {
	godotenv.Load()
	return envConfig{
		DBHost: os.Getenv("DB_HOST"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
		DBPort: os.Getenv("DB_PORT"),
	}
}

// TODO: Replantear la estructura de la conf structs
// para que faciliten las cosas con el archivo toml y
// a la vez que sea lo más idóneo para enviar a las vistas

type Config interface {
	GetFields() map[string]string
}

var (
	confDir  string = getConfDir()
	confFile string = confDir + "/conf.toml"
)

type theme string

const (
	night     theme = "night"
	cupcake   theme = "cupcake"
	halloween theme = "halloween"
	black     theme = "black"
	sunset    theme = "sunset"
)

type UserConfig struct {
	GeneralConf generalConfig
	InfoConf    infoConfig
}

func loadUserConfig() UserConfig {
	return UserConfig{
		GeneralConf: generalConfig{
			CurrentTheme: night,
		},
		InfoConf: infoConfig{
			InfoTiming: time.Second * 2,
		},
	}
}

var UsConf = loadUserConfig()

type infoConfig struct {
	InfoTiming time.Duration `toml:"info_timing"`
}

func (c infoConfig) GetFields() map[string]string {
	m := make(map[string]string)
	m["Info Update Timing"] = c.InfoTiming.String()

	return m
}

type generalConfig struct {
	CurrentTheme theme `toml:"theme"`
}

func (c generalConfig) GetFields() map[string]string {
	m := make(map[string]string)
	m["Theme"] = string(c.CurrentTheme)

	return m
}

func getConfDir() string {
	confDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	confDir = confDir + "/raspi"

	if _, err := os.Stat(confDir); os.IsNotExist(err) {
		err := os.Mkdir(confDir, os.FileMode(0o744))
		if err != nil {
			log.Fatal(err)
		}
	}

	return confDir
}
