package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var JsonConfigs *SettingConfig

func InitJsonConfigs() *SettingConfig {
	JsonConfigs = loadJsonVariables()
	return JsonConfigs
}

type Postgres struct {
	Host    string `json:"host"`
	Port    string `json:"port"`
	User    string `json:"user"`
	DBName  string `json:"DBName"`
	SSLMode string `json:"SSLMode"`
	Token   string `json:"token"`
}

type Todo struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type SettingConfig struct {
	Postgres Postgres `json:"postgres"`
	Todo     Todo     `json:"todo"`
}

func loadJsonVariables() (config *SettingConfig) {
	configFile, err := os.Open("../pkg/config/config.json")
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&config); err != nil {
		fmt.Println("Error parsing config file:", err)
		return
	}
	return config
}
