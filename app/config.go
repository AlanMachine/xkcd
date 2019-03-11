package app

import (
	"encoding/json"
	"os"
)

const configFile = "config.json"

type dbConfig struct {
	DBUser string `json:"db_user"`
	DBPass string `json:"db_pass"`
	DBName string `json:"db_name"`
}

func getConfig() (*dbConfig, error) {
	configPath := os.Getenv("XKCD_CONFIG")

	if configPath == "" {
		configPath = configFile
	}

	f, err := os.OpenFile(configPath, os.O_RDONLY, 0644)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	var conf dbConfig

	if err := json.NewDecoder(f).Decode(&conf); err != nil {
		return nil, err
	}

	return &conf, nil
}
