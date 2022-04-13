package tkapi

import (
	"fmt"
	"time"

	"github.com/didadadida93/lego/pkg/request"
)

type GameSession struct {
	Msid             string
	LobbySession     string
	LobbyCookie      request.Cookie
	GameworldSession string
	GameworldCookie  request.Cookie
	Expires          time.Time
}

func (gs GameSession) GetGameCookie() string {
	return fmt.Sprintf("%s%s;", gs.LobbyCookie.String(), gs.GameworldCookie.String())
}
