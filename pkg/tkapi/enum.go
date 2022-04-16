package tkapi

import (
	"encoding/json"
	"errors"
	"strconv"
)

type resource int

const (
	Wood resource = iota + 1
	Clay
	Iron
	Crop
)

type resources map[resource]int

func (r *resources) UnmarshalJSON(b []byte) error {
	// for the love of god, resources value might be a string
	// so i need to make this function goddamit
	var t map[resource]any
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	f := make(resources)
	for k, v := range t {
		switch t := v.(type) {
		case int:
			f[k] = t
		case float64:
			f[k] = int(t)
		case string:
			if i, err := strconv.Atoi(t); err != nil {
				return err
			} else {
				f[k] = i
			}
		default:
			return errors.New("Unknown type")
		}
	}
	*r = f
	return nil
}
