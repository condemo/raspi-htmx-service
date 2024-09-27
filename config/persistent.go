package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

func SaveConf() error {
	f, err := os.Create(confFile)
	if err != nil {
		return err
	}
	defer f.Close()

	err = toml.NewEncoder(f).Encode(UsConf)
	if err != nil {
		return err
	}

	fmt.Println("Config Saved -> ", UsConf)
	return nil
}

func UpdateConf(c UserConfig) error {
	UsConf = c
	if err := SaveConf(); err != nil {
		return err
	}
	return nil
}