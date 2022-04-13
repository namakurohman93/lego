package login

import (
	"log"
	"net/http"
	"time"

	"github.com/didadadida93/lego/pkg/request"
)

func loginToLobby(e, p string) (c Cookie, s, m string, t time.Time) {
	var token, redirectUrl string
	rc := request.NewRequestConfig()
	getMsid(rc, &m)
	getToken(rc, e, p, &m, &token)
	getRedirectUrl(rc, &m, &token, &redirectUrl)

	rc.Set("url", redirectUrl)
	rc.Set("params", nil)
	rc.Set("body", nil)
	rc.Set("header", nil)
	rc.Set("method", http.MethodGet)
	rc.Set("followRedirect", false)
	res, err := request.Do(rc)
	if err != nil {
		log.Fatal(err)
	}

	// get cookies & session
	c = getCookie(res.Header.Values("set-cookie"))
	s = getSession("gl5SessionKey", res.Header.Values("set-cookie"))
	t = getCookieExp(res.Header.Values("set-cookie"))
	return
}

func getMsid(rc *request.RequestConfig, msid *string) {
	rc.Set("url", "https://mellon-t5.traviangames.com/authentication/login/ajax/form-validate")
	rc.Set("params", nil)
	rc.Set("body", nil)
	rc.Set("header", nil)
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
	rc.Set("header", nil)
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
	rc.Set("header", nil)
	rc.Set("method", http.MethodGet)
	rc.Set("followRedirect", false)

	res, err := request.Do(rc)
	if err != nil {
		log.Fatal(err)
	}

	*r = res.Header.Get("location")
	return
}
