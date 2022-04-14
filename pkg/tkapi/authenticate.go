package tkapi

import (
	"time"

	"github.com/didadadida93/lego/pkg/config"
	"github.com/didadadida93/lego/pkg/login"
)

func GetGameSession(c *config.Config) (g login.GameSession, err error) {
	z := time.Time{}
	if c.GameSession.Msid != "" &&
		c.GameSession.LobbySession != "" &&
		c.GameSession.LobbyCookie != nil &&
		c.GameSession.GameworldSession != "" &&
		c.GameSession.GameworldCookie != nil &&
		c.GameSession.Expires != z {
		return login.GameSession{
			Msid:             c.GameSession.Msid,
			LobbySession:     c.GameSession.LobbySession,
			LobbyCookie:      c.GameSession.LobbyCookie,
			GameworldSession: c.GameSession.GameworldSession,
			GameworldCookie:  c.GameSession.GameworldCookie,
			Expires:          c.GameSession.Expires,
		}, nil
	}
	g, err = Authenticate(c)
	return
}

func Authenticate(c *config.Config) (g login.GameSession, err error) {
	g, err = login.Login(c.Email, c.Password, c.Gameworld)
	if err != nil {
		return
	}
	err = c.UpdateGameSessionConfig(&g)
	return
}
