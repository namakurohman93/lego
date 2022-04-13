package config

import (
	"encoding/json"
	"os"

	"github.com/didadadida93/lego/pkg/login"
)

type Config struct {
	Email       string            `json:"email"`
	Password    string            `json:"password"`
	Gameworld   string            `json:"gameworld"`
	GameSession login.GameSession `json:"gameSession"`
}

func (c *Config) UpdateGameSessionConfig(g *login.GameSession) error {
	c.GameSession = *g
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile("config.json", b, 0644) // -rw-r--r--
	return err
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
