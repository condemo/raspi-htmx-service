package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// TODO: Replantear la estructura de la conf structs
// para que faciliten las cosas con el archivo toml y
// a la vez que sea lo más idóneo para enviar a las vistas

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

var Envs = initConfig()

type envConfig struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

func initConfig() envConfig {
	godotenv.Load()
	return envConfig{
		DBHost: os.Getenv("DB_HOST"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
		DBPort: os.Getenv("DB_PORT"),
	}
}

type uiConfig struct {
	CurrentTheme theme `toml:"theme"`
}

func loadUIConfig() uiConfig {
	return uiConfig{
		CurrentTheme: night,
	}
}

var UIConf = loadUIConfig()

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
