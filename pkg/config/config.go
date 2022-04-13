package config

import (
	"encoding/json"
	"os"

	"github.com/didadadida93/lego/pkg/login"
	"github.com/didadadida93/lego/pkg/tkapi"
)

type Config struct {
	Email       string            `json:"email"`
	Password    string            `json:"password"`
	Gameworld   string            `json:"gameworld"`
	GameSession tkapi.GameSession `json:"gameSession"`
}

func (conf Config) Authenticate() (tkapi.GameSession, error) {
	if conf.GameSession.Msid != "" &&
		conf.GameSession.LobbySession != "" &&
		conf.GameSession.LobbyCookie != nil &&
		conf.GameSession.GameworldSession != "" &&
		conf.GameSession.GameworldCookie != nil {
		return tkapi.GameSession{
			Msid:             conf.GameSession.Msid,
			LobbySession:     conf.GameSession.LobbySession,
			LobbyCookie:      conf.GameSession.LobbyCookie,
			GameworldSession: conf.GameSession.GameworldSession,
			GameworldCookie:  conf.GameSession.GameworldCookie,
		}, nil
	}
	m, s, gs, c, gc := login.Login(conf.Email, conf.Password, conf.Gameworld)
	gameSession := tkapi.GameSession{
		Msid:             m,
		LobbySession:     s,
		LobbyCookie:      c,
		GameworldSession: gs,
		GameworldCookie:  gc,
	}
	err := saveGameSession(&conf, &gameSession)
	return gameSession, err
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

func saveGameSession(conf *Config, gs *tkapi.GameSession) error {
	conf.GameSession = *gs
	b, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile("config.json", b, 0644) // -rw-r--r--
	return err
}
