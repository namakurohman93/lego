package tkapi

type playerId string

type player struct {
	Active       string       `json:"active"`
	Country      string       `json:"country"`
	KingId       int          `json:"kingId"`
	KingStatus   bool         `json:"kingStatus"`
	KingdomId    string       `json:"kingdomId"`
	KingdomRole  string       `json:"kingdomRole"`
	Name         string       `json:"name"`
	SpawnedOnMap coordinateId `json:"spawnedOnMap,string"`
	TribeId      string       `json:"tribeId"`
}

type players map[playerId]player
