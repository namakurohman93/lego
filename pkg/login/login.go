package login

import (
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"
)

type Cookie map[string]string

func (c Cookie) String() (s string) {
	for k, v := range c {
		s += fmt.Sprintf("%s=%v; ", k, v)
	}
	return
}

func Login(e, p, gw string) (string, string, string, Cookie, Cookie, time.Time) {
	c, s, m, t := loginToLobby(e, p)
	gc, gs := loginToGameworld(c, s, m, gw)
	return m, s, gs, c, gc, t
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

func getCookieExp(cookies []string) time.Time {
	s := strings.Split(strings.Split(cookies[0], ";")[1], "=")[1]
	t, err := time.Parse(time.RFC1123, strings.ReplaceAll(s, "-", " "))
	if err != nil {
		log.Fatal(err)
	}
	return t
}
