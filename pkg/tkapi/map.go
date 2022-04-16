package tkapi

import (
	"encoding/json"
	"github.com/didadadida93/lego/pkg/request"
	"net/http"
	"strings"
)

type coordinateId string

type cell interface {
	IsRural() bool
	IsOasis() bool
	IsAbandonedValley() bool
	IsRobberHideout() bool
}

// type cells []json.RawMessage
type cells []cell

func (cs *cells) GetOases() *[]oasis {
	var r []oasis
	for _, e := range *cs {
		switch v := e.(type) {
		case oasis:
			r = append(r, v)
		default:
			continue
		}
	}
	return &r
}

func (cs *cells) GetAbandonedValleys() *[]abandonedValley {
	var r []abandonedValley
	for _, e := range *cs {
		switch v := e.(type) {
		case abandonedValley:
			r = append(r, v)
		default:
			continue
		}
	}
	return &r
}

func (cs *cells) GetRurals() *[]rural {
	var r []rural
	for _, e := range *cs {
		switch v := e.(type) {
		case rural:
			r = append(r, v)
		default:
			continue
		}
	}
	return &r
}

func (cs *cells) GetRobberHideouts() *[]robberHideout {
	var r []robberHideout
	for _, e := range *cs {
		switch v := e.(type) {
		case robberHideout:
			r = append(r, v)
		default:
			continue
		}
	}
	return &r
}

type region map[coordinateId][]json.RawMessage

type mapData struct {
	Kingdom any    `json:"kingdom"`
	Player  any    `json:"player"`
	Region  region `json:"region"`
	Report  any    `json:"report"`
}

type cellFiller struct {
	Id        string `json:"id"`
	Landscape string `json:"landscape"`
	Owner     string `json:"owner"`
}

func (c cellFiller) IsRural() bool           { return false }
func (c cellFiller) IsOasis() bool           { return false }
func (c cellFiller) IsAbandonedValley() bool { return false }
func (c cellFiller) IsRobberHideout() bool   { return false }

type oasis struct {
	Id        string    `json:"id"`
	Landscape string    `json:"landscape"`
	Owner     string    `json:"owner"`
	Oasis     oasisData `json:"oasis"`
}
type oasisData struct {
	Bonus       resources `json:"bonus"`
	KingId      string    `json:"kingId"`
	KingdomId   string    `json:"kingdomId"`
	OasisStatus string    `json:"oasisStatus"`
	Type        string    `json:"type"`
	Units       []any     `json:"units"`
}

func (o oasis) IsRural() bool           { return false }
func (o oasis) IsOasis() bool           { return true }
func (o oasis) IsAbandonedValley() bool { return false }
func (o oasis) IsRobberHideout() bool   { return false }

type rural struct {
	Id        string      `json:"id"`
	Landscape string      `json:"landscape"`
	Owner     string      `json:"owner"`
	PlayerId  string      `json:"playerId"`
	ResType   string      `json:"resType"`
	Village   villageData `json:"village"`
}
type villageData struct {
	HasActiveTreasury bool   `json:"hasActiveTreasury"`
	Name              string `json:"name"`
	Population        int    `json:"population,string"`
	Type              string `json:"type"`
	VillageId         string `json:"villageId"`
}

func (r rural) IsRural() bool           { return true }
func (r rural) IsOasis() bool           { return false }
func (r rural) IsAbandonedValley() bool { return false }
func (r rural) IsRobberHideout() bool   { return false }

type abandonedValley struct {
	Id        string `json:"id"`
	Landscape string `json:"landscape"`
	Owner     string `json:"owner"`
	ResType   string `json:"resType"`
}

func (a abandonedValley) IsRural() bool           { return false }
func (a abandonedValley) IsOasis() bool           { return false }
func (a abandonedValley) IsAbandonedValley() bool { return true }
func (a abandonedValley) IsRobberHideout() bool   { return false }

type robberHideout struct {
	Id        string            `json:"id"`
	Landscape string            `json:"landscape"`
	Owner     string            `json:"owner"`
	PlayerId  int               `json:"playerId"`
	Village   robberHideoutData `json:"village"`
}
type robberHideoutData struct {
	Name       string `json:"name"`
	Population int    `json:"population"`
	Type       int    `json:"type"`
	VillageId  string `json:"villageId"` // be aware that this village id is unusual
}

func (r robberHideout) IsRural() bool           { return false }
func (r robberHideout) IsOasis() bool           { return false }
func (r robberHideout) IsAbandonedValley() bool { return false }
func (r robberHideout) IsRobberHideout() bool   { return true }

func coordinateToCellId(x, y int) int {
	return (536887296 + x) + (y * 32768)
}

func regIds() (r []int) {
	for x := -13; x < 14; x++ {
		for y := -13; y < 14; y++ {
			r = append(r, coordinateToCellId(x, y))
		}
	}
	return
}

func (gd *GameDriver) RequestMap() (*cells, error) {
	if err := checkExpired(gd); err != nil {
		return nil, err
	}

	ids := regIds()
	c, a := "map", "getByRegionIds"
	url := gd.GetUrl(c, a)
	rc := request.NewRequestConfig()
	rc.Set("url", url)
	rc.Set("params", nil)
	rc.Set("body", &payload{
		Action:     a,
		Controller: c,
		Session:    gd.GameSession.GameworldSession,
		Params: getMapParams{
			RegionIdCollection: map[string][]int{"1": ids},
		},
	})
	rc.Set("header", request.Header{
		"Cookie": gd.GameSession.GetGameCookie(),
	})
	rc.Set("method", http.MethodPost)
	rc.Set("followRedirect", false)

	res, err := request.Do(rc)
	if err != nil {
		return nil, err
	}
	var r response
	if err := json.Unmarshal([]byte(res.Body), &r); err != nil {
		return nil, err
	}
	var md mapData
	if err := shiftType(r.Response["1"], &md); err != nil {
		return nil, err
	}
	var cs cells
	for _, v := range md.Region {
		for _, c := range v {
			t := string(c)
			if strings.Contains(t, "oasis") {
				var f oasis
				if err := json.Unmarshal(c, &f); err != nil {
					return nil, err
				}
				cs = append(cs, f)
			} else if strings.Contains(t, "resType") &&
				!strings.Contains(t, "playerId") &&
				!strings.Contains(t, "village") {
				var f abandonedValley
				if err := json.Unmarshal(c, &f); err != nil {
					return nil, err
				}
				cs = append(cs, f)
			} else if strings.Contains(t, "playerId") &&
				strings.Contains(t, "resType") &&
				strings.Contains(t, "village") {
				var f rural
				if err := json.Unmarshal(c, &f); err != nil {
					return nil, err
				}
				cs = append(cs, f)
			} else if strings.Contains(t, "playerId") &&
				strings.Contains(t, "village") &&
				!strings.Contains(t, "resType") {
				var f robberHideout
				if err := json.Unmarshal(c, &f); err != nil {
					return nil, err
				}
				cs = append(cs, f)
			} else {
				var f cellFiller
				if err := json.Unmarshal(c, &f); err != nil {
					return nil, err
				}
				cs = append(cs, f)
			}
		}
	}
	return &cs, nil
}
