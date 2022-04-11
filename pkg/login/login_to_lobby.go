package login

import (
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/didadadida93/lego/pkg/request"
)

type LobbyCookies map[string]string

func loginToLobby(e, p string) (c LobbyCookies, s string) {
	var msid, token, redirectUrl string
	rc := request.NewRequestConfig()
	getMsid(rc, &msid)
	getToken(rc, e, p, &msid, &token)
	getRedirectUrl(rc, &msid, &token, &redirectUrl)

	rc.Set("url", redirectUrl)
	rc.Set("params", nil)
	rc.Set("body", nil)
	rc.Set("method", http.MethodGet)
	rc.Set("followRedirect", false)
	res, err := request.Do(rc)
	if err != nil {
		log.Fatal(err)
	}

	// get cookies & session
	c = getLobbyCookie(res.Header.Values("set-cookie"))
	s = getLobbySession(res.Header.Values("set-cookie"))
	return
}

func getMsid(rc *request.RequestConfig, msid *string) {
	rc.Set("url", "https://mellon-t5.traviangames.com/authentication/login/ajax/form-validate")
	rc.Set("params", nil)
	rc.Set("body", nil)
	rc.Set("method", http.MethodGet)
	rc.Set("followRedirect", false)

	res, err := request.Do(rc)
	if err != nil {
		log.Fatal(err)
	}
	execRegexp(`msid=([\w]*)&msname`, res.Body, msid)
	return
}

func getToken(rc *request.RequestConfig, e, p string, m, t *string) {
	rc.Set("url", "https://mellon-t5.traviangames.com/authentication/login/ajax/form-validate")
	rc.Set("params", request.UrlParams{
		"msid":   *m,
		"msname": "msid",
	})
	rc.Set("body", &request.FormBody{
		"email":    e,
		"password": p,
	})
	rc.Set("method", http.MethodPost)
	rc.Set("followRedirect", true)

	res, err := request.Do(rc)
	if err != nil {
		log.Fatal(err)
	}
	execRegexp(`token=([\w]*)&msid`, res.Body, t)
	return
}

func getRedirectUrl(rc *request.RequestConfig, m, t, r *string) {
	rc.Set("url", "http://lobby.kingdoms.com/api/login.php")
	rc.Set("params", request.UrlParams{
		"msid":   *m,
		"token":  *t,
		"msname": "msid",
	})
	rc.Set("body", nil)
	rc.Set("method", http.MethodGet)
	rc.Set("followRedirect", false)

	res, err := request.Do(rc)
	if err != nil {
		log.Fatal(err)
	}

	*r = res.Header.Get("location")
	return
}

func getLobbyCookie(cookies []string) LobbyCookies {
	result := make(map[string]string)

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

func getLobbySession(cookies []string) string {
	t := getLobbyCookie(cookies)
	s := t["gl5SessionKey"]
	v := strings.Split(s, `"`)[3]
	return v
}

func execRegexp(r, s string, t *string) {
	re := regexp.MustCompile(r)
	*t = re.FindStringSubmatch(s)[1]
	return
}
