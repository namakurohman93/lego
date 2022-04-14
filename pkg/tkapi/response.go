package tkapi

import "encoding/json"

type response struct {
	Cache    []cache `json:"cache"`
	Error    any     `json:"error"`
	Response []any   `json:"response"`
	SerialNo int     `json:"serialNo"`
	Time     int64   `json:"time"`
}

type cache struct {
	Name string    `json:"name"`
	Data cacheData `json:"data"`
}

type cacheData struct {
	Cache     []map[string]any `json:"cache"`
	Operation int              `json:"operation"`
}

func processCacheData(d, r any) error {
	c, err := json.Marshal(d)
	if err != nil {
		return err
	}
	err = json.Unmarshal(c, r)
	if err != nil {
		return err
	}
	return nil
}
