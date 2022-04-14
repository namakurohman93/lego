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
	var t map[resource]any
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	f := make(resources)
	for k, v := range t {
		switch v.(type) {
		case int:
			f[k] = v.(int)
		case float64:
			f[k] = int(v.(float64))
		case string:
			if i, err := strconv.Atoi(v.(string)); err != nil {
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
