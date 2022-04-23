package tkapi

import "strings"

type kingdomId string

type kingdom struct {
	Tag string `json:"tag"`
}

type kingdoms map[kingdomId]kingdom

func (ks kingdoms) GetByName(name string) (r kingdom, ok bool) {
	for _, k := range ks {
		if strings.ToLower(k.Tag) == strings.ToLower(name) {
			ok = true
			r = k
			return
		}
	}
	return
}
