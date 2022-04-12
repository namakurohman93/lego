package login

import (
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
)

type Cookie map[string]string

func (c Cookie) String() (s string) {
	for k, v := range c {
		s += fmt.Sprintf("%s=%v; ", k, v)
	}
	return
}

func Login() {
	e, p, gw := "email@email.com", "password", "comx"
	c, s, m := loginToLobby(e, p)
	gc, gs := loginToGameworld(c, s, m, gw)
	log.Println(c)
	log.Println(s)
	log.Println(m)
	log.Println(gc)
	log.Println(gs)
}

func execRegexp(r, s string, t *string) {
	re := regexp.MustCompile(r)
	*t = re.FindStringSubmatch(s)[1]
	return
}

func getCookie(cookies []string) Cookie {
	result := make(Cookie)

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
