package config

import (
	"encoding/json"
	"os"
	"time"

	"github.com/didadadida93/lego/pkg/login"
)

type Config struct {
	Email       string            `json:"email"`
	Password    string            `json:"password"`
	Gameworld   string            `json:"gameworld"`
	GameSession login.GameSession `json:"gameSession"`
}

func (conf Config) Authenticate() (login.GameSession, error) {
	z := time.Time{}
	if conf.GameSession.Msid != "" &&
		conf.GameSession.LobbySession != "" &&
		conf.GameSession.LobbyCookie != nil &&
		conf.GameSession.GameworldSession != "" &&
		conf.GameSession.GameworldCookie != nil &&
		conf.GameSession.Expires != z {

		return login.GameSession{
			Msid:             conf.GameSession.Msid,
			LobbySession:     conf.GameSession.LobbySession,
			LobbyCookie:      conf.GameSession.LobbyCookie,
			GameworldSession: conf.GameSession.GameworldSession,
			GameworldCookie:  conf.GameSession.GameworldCookie,
			Expires:          conf.GameSession.Expires,
		}, nil
	}

	gameSession := login.Login(conf.Email, conf.Password, conf.Gameworld)
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

func saveGameSession(conf *Config, gs *login.GameSession) error {
	conf.GameSession = *gs
	b, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile("config.json", b, 0644) // -rw-r--r--
	return err
}
