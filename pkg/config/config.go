package config

import (
	"encoding/json"
	"os"

	"github.com/didadadida93/lego/pkg/login"
	"github.com/didadadida93/lego/pkg/tkapi"
)

type Config struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Gameworld string `json:"gameworld"`
}

func (conf Config) Authenticate() tkapi.GameSession {
	m, s, gs, c, gc := login.Login(conf.Email, conf.Password, conf.Gameworld)
	return tkapi.GameSession{
		Msid:             m,
		LobbySession:     s,
		LobbyCookie:      c,
		GameworldSession: gs,
		GameworldCookie:  gc,
	}
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
