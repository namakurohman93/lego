package login

import (
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
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

func Login(e, p, gw string) (g GameSession, err error) {
	c, s, m, t, err := loginToLobby(e, p)
	if err != nil {
		return
	}

	gc, gs, err := loginToGameworld(c, s, m, gw)
	if err != nil {
		return
	}

	g.Msid = m
	g.LobbySession = s
	g.LobbyCookie = c
	g.GameworldSession = gs
	g.GameworldCookie = gc
	g.Expires = t
	return
}

func execRegexp(r, s string, t *string) {
	re := regexp.MustCompile(r)
	*t = re.FindStringSubmatch(s)[1]

	return
}

func getCookie(cookies []string) request.Cookie {
	result := make(request.Cookie)
	for _, cookie := range cookies {
		c := strings.Split(cookie, ";")[0]
		t := strings.Split(c, "=")
		k, v := t[0], t[1]
		v1, err := url.QueryUnescape(v)
		if err != nil {
			log.Fatal(err)
		}
		if _, ok := result[k]; !ok {
			result[k] = v1
		}
	}

	return result
}

func getSession(key string, cookies []string) string {
	t := getCookie(cookies)
	s := t[key]
	v := strings.Split(s, `"`)[3]

	return v
}

func getCookieExp(cookies []string) (t time.Time, err error) {
	s := strings.Split(strings.Split(cookies[0], ";")[1], "=")[1]
	t, err = time.Parse(time.RFC1123, strings.ReplaceAll(s, "-", " "))
	return
}
