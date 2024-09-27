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

var (
	confDir  string = getConfDir()
	confFile string = confDir + "/conf.toml"
)

type Theme string

const (
	night     Theme = "night"
	cupcake   Theme = "cupcake"
	halloween Theme = "halloween"
	black     Theme = "black"
	sunset    Theme = "sunset"
)

type UserConfig struct {
	GeneralConf GeneralConfig
	InfoConf    InfoConfig
}

func loadUserConfig() UserConfig {
	return UserConfig{
		GeneralConf: GeneralConfig{
			CurrentTheme: night,
		},
		InfoConf: InfoConfig{
			InfoTick: time.Second * 2,
		},
	}
}

var UsConf = loadUserConfig()

type InfoConfig struct {
	InfoTick time.Duration `toml:"info_tick"`
}

type GeneralConfig struct {
	CurrentTheme Theme `toml:"theme"`
}

func (c GeneralConfig) GetThemeList() []Theme {
	return []Theme{
		night, cupcake, halloween,
		black, sunset,
	}
}

// TODO: Implementar un sistema para leer `conf.toml`
// y cargar la config en `UserConfig` struct
// Decidir si no hacer global la variable `UsConf`

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
