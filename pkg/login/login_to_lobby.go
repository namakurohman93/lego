package login

import (
	"net/http"
	"time"

	"github.com/didadadida93/lego/pkg/request"
)

func loginToLobby(e, p string) (c request.Cookie, s, m string, t time.Time, err error) {
	var token, redirectUrl string
	rc := request.NewRequestConfig()
	err = getMsid(rc, &m)
	if err != nil {
		return
	}
	err = getToken(rc, e, p, &m, &token)
	if err != nil {
		return
	}
	err = getRedirectUrl(rc, &m, &token, &redirectUrl)
	if err != nil {
		return
	}

	rc.Set("url", redirectUrl)
	rc.Set("params", nil)
	rc.Set("body", nil)
	rc.Set("header", nil)
	rc.Set("method", http.MethodGet)
	rc.Set("followRedirect", false)

	res, err := request.Do(rc)
	if err != nil {
		return
	}

	// get cookies & session
	c = getCookie(res.Header.Values("set-cookie"))
	s = getSession("gl5SessionKey", res.Header.Values("set-cookie"))
	t, err = getCookieExp(res.Header.Values("set-cookie"))
	return
}

func getMsid(rc *request.RequestConfig, msid *string) error {
	rc.Set("url", "https://mellon-t5.traviangames.com/authentication/login/ajax/form-validate")
	rc.Set("params", nil)
	rc.Set("body", nil)
	rc.Set("header", nil)
	rc.Set("method", http.MethodGet)
	rc.Set("followRedirect", false)

	res, err := request.Do(rc)
	if err != nil {
		return err
	}
	execRegexp(`msid=([\w]*)&msname`, res.Body, msid)
	return nil
}

func getToken(rc *request.RequestConfig, e, p string, m, t *string) error {
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
		return err
	}
	execRegexp(`token=([\w]*)&msid`, res.Body, t)
	return nil
}

func getRedirectUrl(rc *request.RequestConfig, m, t, r *string) error {
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
		return err
	}

	*r = res.Header.Get("location")
	return nil
}
