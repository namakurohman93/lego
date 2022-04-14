package tkapi

import (
	"fmt"
	"time"

	"github.com/didadadida93/lego/pkg/config"
	"github.com/didadadida93/lego/pkg/login"
)

var gameworldUrl string = "https://%s.kingdoms.com/api/?c=%s&a=%s&t%v"

type GameDriver struct {
	Config      *config.Config
	GameSession login.GameSession
}

func (gd *GameDriver) GetUrl(c, a string) string {
	return fmt.Sprintf(gameworldUrl, gd.Config.Gameworld, c, a, time.Now().Unix())
}

func NewDriver(c *config.Config) (*GameDriver, error) {
	z := time.Time{}
	if c.GameSession.Msid != "" &&
		c.GameSession.LobbySession != "" &&
		c.GameSession.LobbyCookie != nil &&
		c.GameSession.GameworldSession != "" &&
		c.GameSession.GameworldCookie != nil &&
		c.GameSession.Expires != z {
		return &GameDriver{
			Config: c,
			GameSession: login.GameSession{
				Msid:             c.GameSession.Msid,
				LobbySession:     c.GameSession.LobbySession,
				LobbyCookie:      c.GameSession.LobbyCookie,
				GameworldSession: c.GameSession.GameworldSession,
				GameworldCookie:  c.GameSession.GameworldCookie,
				Expires:          c.GameSession.Expires,
			},
		}, nil
	}
	g, err := authenticate(c)
	return &GameDriver{
		Config:      c,
		GameSession: g,
	}, err
}

func checkExpired(gd *GameDriver) error {
	y := gd.GameSession.Expires.AddDate(0, 0, -1)
	n := time.Now()
	if yes := y.Before(n); yes {
		err := gd.ReAuthenticate()
		return err
	}
	return nil
}
