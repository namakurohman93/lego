package tkapi

import "math"

type coordinateId int

func (c coordinateId) ParseCoordinate() coordinates {
	return coordinates{
		X: int((c % 32768) - 16384),
		Y: int(math.Floor(float64(c)/32768) - 16384),
	}
}

type coordinates struct {
	X int `json:"x,string"`
	Y int `json:"y,string"`
}
