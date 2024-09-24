package config

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

func SaveConf() {
	f, err := os.Create(confFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = toml.NewEncoder(f).Encode(UsConf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Config Saved")
}
