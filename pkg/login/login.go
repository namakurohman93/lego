package login

import (
	"github.com/didadadida93/lego/pkg/request"
	"log"
	"net/url"
	"regexp"
	"strings"
)

func Login() {
	e, p := "email@email.com", "password"
	c, s := loginToLobby(e, p)
	log.Println(c)
	log.Println(s)
}

func loginToLobby(e, p string) (c map[string]string, s string) {
	// get msid
	url := "https://mellon-t5.traviangames.com/authentication/login/ajax/form-validate"
	res, err := request.Get(url, nil, false)
	if err != nil {
		log.Fatal(err)
	}
	msid := execRegexp(`msid=([\w]*)&msname`, res.Body)

	// get lobby token
	params := request.UrlParams{
		"msid":   msid,
		"msname": "msid",
	}
	credential := &request.FormBody{
		"email":    e,
		"password": p,
	}
	res, err = request.PostForm(url, params, credential, true)
	if err != nil {
		log.Fatal(err)
	}
	token := execRegexp(`token=([\w]*)&msid`, res.Body)

	// get redirect URI
	url = "http://lobby.kingdoms.com/api/login.php"
	params = request.UrlParams{
		"msid":   msid,
		"token":  token,
		"msname": "msid",
	}
	res, err = request.Get(url, params, false)
	if err != nil {
		log.Fatal(err)
	}

	res, err = request.Get(res.Header.Get("location"), nil, false)
	if err != nil {
		log.Fatal(err)
	}

	// get cookies & session
	c = getLobbyCookie(res.Header.Values("set-cookie"))
	s = getLobbySession(res.Header.Values("set-cookie"))
	return
}

func getLobbyCookie(cookies []string) map[string]string {
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

func execRegexp(r, s string) string {
	re := regexp.MustCompile(r)
	return re.FindStringSubmatch(s)[1]
}
