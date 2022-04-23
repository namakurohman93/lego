package tkapi

import "encoding/json"

type response struct {
	Cache    []cache         `json:"cache"`
	Error    errorResponse   `json:"error,omitempty"`
	Response json.RawMessage `json:"response"`
	SerialNo int             `json:"serialNo"`
	Time     int64           `json:"time"`
}

type cache struct {
	Name string    `json:"name"`
	Data cacheData `json:"data"`
}

type cacheData struct {
	Cache     []map[string]any `json:"cache"`
	Operation int              `json:"operation"`
}

type errorResponse struct {
	Type    string `json:"type"`
	Number  int    `json:"number"`
	Message string `json:"message"`
}

func shiftType(d, r any) error {
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
