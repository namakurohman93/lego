package tkapi

import (
	"github.com/didadadida93/lego/pkg/config"
	"github.com/didadadida93/lego/pkg/login"
)

func Authenticate(c *config.Config) (g login.GameSession, err error) {
	g, err = login.Login(c.Email, c.Password, c.Gameworld)
	if err != nil {
		return
	}
	err = c.UpdateGameSessionConfig(&g)
	return
}

func (gd *GameDriver) ReAuthenticate() error {
	_, err := Authenticate(gd.Config)
	return err
}
