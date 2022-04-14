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

func RequestOwnVillage(gw, gs, cookie string) (vs []village, err error) {
	controller, action := "cache", "get"
	url := fmt.Sprintf(gameworldUrl, gw, controller,
		action, time.Now().Unix())
	rc := request.NewRequestConfig()
	rc.Set("url", url)
	rc.Set("params", nil)
	rc.Set("body", &payload{
		Action:     action,
		Controller: controller,
		Session:    gs,
		Params: cacheParams{
			Names: []string{"Collection:Village:own"},
		},
	})
	rc.Set("header", request.Header{
		"Cookie": cookie,
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
