package tkapi

import (
	"encoding/json"
	"net/http"

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
	Production       resources   `json:"production"`
	Storage          resources   `json:"storage"`
	StorageCapacity  resources   `json:"storageCapacity"`
}

type coordinates struct {
	X int `json:"x,string"`
	Y int `json:"y,string"`
}

func (gd *GameDriver) RequestOwnVillage() (vs []village, err error) {
	err = checkExpired(gd)
	if err != nil {
		return
	}
	c, a := "cache", "get"
	url := gd.GetUrl(c, a)
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
		err = shiftType(cv["data"], &v)
		if err != nil {
			return vs, err
		}
		vs = append(vs, v)
	}
	return
}
