package tkapi

import (
	"time"

	"github.com/didadadida93/lego/pkg/config"
	"github.com/didadadida93/lego/pkg/login"
)

var gameworldUrl string = "https://%s.kingdoms.com/api/?c=%s&a=%s&t%v"

type GameDriver struct {
	Config      *config.Config
	GameSession login.GameSession
}

func NewDriver(c *config.Config) (*GameDriver, error) {
	z := time.Time{}
	if c.GameSession.Msid != "" &&
		c.GameSession.LobbySession != "" &&
		c.GameSession.LobbyCookie != nil &&
		c.GameSession.GameworldSession != "" &&
		c.GameSession.GameworldCookie != nil &&
		c.GameSession.Expires != z {
		gs := login.GameSession{
			Msid:             c.GameSession.Msid,
			LobbySession:     c.GameSession.LobbySession,
			LobbyCookie:      c.GameSession.LobbyCookie,
			GameworldSession: c.GameSession.GameworldSession,
			GameworldCookie:  c.GameSession.GameworldCookie,
			Expires:          c.GameSession.Expires,
		}
		return &GameDriver{
			Config:      c,
			GameSession: gs,
		}, nil
	}
	g, err := Authenticate(c)
	return &GameDriver{
		Config:      c,
		GameSession: g,
	}, err
}

func getExpired(gd *GameDriver) bool {
	y := gd.GameSession.Expires.AddDate(0, 0, -1)
	n := time.Now()
	return y.Before(n)
}
