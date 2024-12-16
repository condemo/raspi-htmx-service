package config

import (
	"log"
	"os"
	"path/filepath"
)

type LoggerConfig struct {
	LogFilePath string
}

func NewLoggerConfig() *LoggerConfig {
	lc := &LoggerConfig{}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		// TODO: Mejorar error handling
		log.Fatal(err)
	}

	logDir := filepath.Join(homeDir, "/.local/share/raspi")
	logFile := filepath.Join(logDir, "errors")

	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.Mkdir(logDir, os.FileMode(0o744))
		if err != nil {
			log.Fatal(err)
		}

		f, err := os.Create(logFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

	}

	lc.LogFilePath = logFile

	return lc
}

func (c *LoggerConfig) SaveLog(msg string) error {
	f, err := os.OpenFile(c.LogFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(msg + "\n"); err != nil {
		return err
	}

	return nil
}

func (c *LoggerConfig) Clean() error {
	if err := os.Truncate(c.LogFilePath, 0); err != nil {
		return err
	}
	return nil
}
