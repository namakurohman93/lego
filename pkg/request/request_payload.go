package request

import (
	"bytes"
	"encoding/json"
	"io"
	"net/url"
	"strings"
)

var formContentType string = "application/x-www-form-urlencoded"
var jsonContentType string = "application/json"

type IPayload interface {
	GetBody() (io.Reader, error)
	GetHeader() Header
}

type FormBody map[string]string

type TKPayload struct {
	Action     string   `json:"action"`
	Controller string   `json:"controller"`
	Session    string   `json:"session"`
	Params     TKParams `json:"params"`
}

type TKParams struct {
	Names []string `json:"names"`
}

func (f *FormBody) GetBody() (io.Reader, error) {
	d := url.Values{}
	for k, v := range *f {
		d.Set(k, v)
	}
	return strings.NewReader(d.Encode()), nil
}

func (f *FormBody) GetHeader() Header {
	return Header{"Content-Type": formContentType}
}

func (p *TKPayload) GetBody() (io.Reader, error) {
	body, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(body), nil
}

func (p *TKPayload) GetHeader() Header {
	return Header{"Content-Type": jsonContentType}
}
