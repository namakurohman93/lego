package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Gameworld string `json:"gameworld"`
}

func GetConfig() (c Config, err error) {
	b, err := os.ReadFile("config.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &c)
	if err != nil {
		return
	}
	return
}
