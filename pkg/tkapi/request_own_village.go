package tkapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/didadadida93/lego/pkg/request"
)

type village struct {
	VillageId        string      `json:"villageId"`
	PlayerId         string      `json:"playerId"`
	Name             string      `json:"name"`
	TribeId          int         `json:"tribeId,string"`
	BelongsToKing    string      `json:"belongsToKing"`
	BelongsToKingdom int         `json:"belongsToKingdom"`
	Type             int         `json:"type,string"`
	Population       int         `json:"population,string"`
	Coordinates      coordinates `json:"coordinates"`
	IsMainVillage    bool        `json:"isMainVillage"`
	IsTown           bool        `json:"isTown"`
	CulturePoints    float32     `json:"culturePoints"`
}

type coordinates struct {
	X int `json:"x,string"`
	Y int `json:"y,string"`
}

func (gd *GameDriver) RequestOwnVillage() (vs []village, err error) {
	if expired := getExpired(gd); expired {
		// need to re authenticate
		// writing session to file again
		// update current gd session
		err = gd.ReAuthenticate()
		if err != nil {
			return vs, err
		}
	}
	c, a := "cache", "get"
	url := fmt.Sprintf(gameworldUrl, gd.Config.Gameworld, c, a, time.Now().Unix())
	rc := request.NewRequestConfig()
	rc.Set("url", url)
	rc.Set("params", nil)
	rc.Set("body", &payload{
		Action:     a,
		Controller: c,
		Session:    gd.GameSession.GameworldSession,
		Params: cacheParams{
			Names: []string{"Collection:Village:own"},
		},
	})
	rc.Set("header", request.Header{
		"Cookie": gd.GameSession.GetGameCookie(),
	})
	rc.Set("method", http.MethodPost)
	rc.Set("followRedirect", false)

	res, err := request.Do(rc)
	if err != nil {
		return
	}
	var resp response
	err = json.Unmarshal([]byte(res.Body), &resp)
	if err != nil {
		return
	}

	for _, cv := range resp.Cache[0].Data.Cache {
		var v village
		err = processCacheData(cv["data"], &v)
		if err != nil {
			return vs, err
		}
		vs = append(vs, v)
	}
	return
}
