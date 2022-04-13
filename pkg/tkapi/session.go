package tkapi

import (
	"fmt"
	"time"

	"github.com/didadadida93/lego/pkg/login"
)

type GameSession struct {
	Msid             string
	LobbySession     string
	LobbyCookie      login.Cookie
	GameworldSession string
	GameworldCookie  login.Cookie
	Expires          time.Time
}

func (gs GameSession) GetGameCookie() string {
	return fmt.Sprintf("%s%s;", gs.LobbyCookie.String(), gs.GameworldCookie.String())
}
