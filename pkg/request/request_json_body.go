package request

import (
	"bytes"
	"encoding/json"
	"io"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *User) GetBody() (io.Reader, error) {
	body, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(body), nil
}

type TKPayload struct {
	Action     string   `json:"action"`
	Controller string   `json:"controller"`
	Session    string   `json:"session"`
	Params     TKParams `json:"params"`
}

type TKParams struct {
	Names []string `json:"names"`
}

func (p *TKPayload) GetBody() (io.Reader, error) {
	body, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(body), nil
}
