package login

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/didadadida93/lego/pkg/request"
)

type tkApiResponse struct {
	Time     int64         `json:"time,string"`
	SerialNo int           `json:"serialNo"`
	Cache    []cacheStruct `json:"cache"`
	Response any           `json:"response"`
	Error    any           `json:"error"`
}

// --- starting from cache, it already different
type cacheStruct struct {
	Data cacheData `json:"data"`
	Name string    `json:"name"`
}

// --- or maybe starting from here, idk
type cacheData struct {
	Cache     []avatarCache `json:"cache"`
	Operation int           `json:"operation"`
}

type avatarCache struct {
	Data avatarStruct `json:"data"`
	Name string       `json:"name"`
}

type avatarStruct struct {
	AccoutnName           string `json:"accountName"`
	AvatarIdentifier      string `json:"avatarIdentifier"`
	AvatarName            string `json:"avatarName"`
	BanPaymentProvider    string `json:"banPaymentProvider"`
	BanReason             string `json:"banReason"`
	ConsumersId           string `json:"consumersId"`
	Country               string `json:"country"`
	IsBanned              bool   `json:"isBanned"`
	IsSuspended           bool   `json:"isSuspended"`
	Limitation            string `json:"limitation"`
	SuspensionTime        string `json:"suspensionTime"`
	UserAccountIdentifier string `json:"userAccountIdentifier"`
	WorldName             string `json:"worldName"`
}

// --- ended here

func loginToGameworld(c request.Cookie, s, m, gw string) (gc request.Cookie, gs string, err error) {
	var token, gwId string
	rc := request.NewRequestConfig()
	err = getGameworldId(rc, s, gw, &gwId)
	if err != nil {
		return
	}
	err = getGameworldToken(rc, m, gwId, &token)
	if err != nil {
		return
	}

	url := fmt.Sprintf("https://%s.kingdoms.com/api/login.php", gw)
	rc.Set("url", url)
	rc.Set("params", request.UrlParams{
		"msid":   m,
		"token":  token,
		"msname": "msid",
	})
	rc.Set("body", nil)
	rc.Set("header", nil)
	rc.Set("method", http.MethodGet)
	rc.Set("followRedirect", false)

	res, err := request.Do(rc)
	if err != nil {
		return
	}

	// get cookie & session
	gc = getCookie(res.Header.Values("set-cookie"))
	gs = getSession("t5SessionKey", res.Header.Values("set-cookie"))
	return
}

func getGameworldId(rc *request.RequestConfig, s, gw string, gwId *string) error {
	rc.Set("url", "https://lobby.kingdoms.com/api/index.php")
	rc.Set("params", nil)
	rc.Set("body", &request.TKPayload{
		Action:     "get",
		Controller: "cache",
		Session:    s,
		Params: request.TKParams{
			Names: []string{"Collection:Avatar:"},
		},
	})
	rc.Set("header", nil)
	rc.Set("method", http.MethodPost)
	rc.Set("followRedirect", true)

	res, err := request.Do(rc)
	if err != nil {
		return err
	}
	var resp tkApiResponse
	err = json.Unmarshal([]byte(res.Body), &resp)
	if err != nil {
		return err
	}
	for _, avatar := range resp.Cache[0].Data.Cache {
		if strings.ToLower(avatar.Data.WorldName) == strings.ToLower(gw) {
			*gwId = avatar.Data.ConsumersId
			break
		}
	}
	return nil
}

func getGameworldToken(rc *request.RequestConfig, m, gwId string, t *string) error {
	url := fmt.Sprintf(
		"https://mellon-t5.traviangames.com/game-world/join/gameWorldId/%s",
		gwId)
	rc.Set("url", url)
	rc.Set("params", request.UrlParams{
		"msid":   m,
		"msname": "msid",
	})
	rc.Set("body", nil)
	rc.Set("header", nil)
	rc.Set("method", http.MethodGet)
	rc.Set("followRedirect", true)

	res, err := request.Do(rc)
	if err != nil {
		return err
	}
	execRegexp(`token=([\w]*)&msid`, res.Body, t)
	return nil
}
