package tkapi

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/didadadida93/lego/pkg/request"
)

var contentType = "application/json"

type payload struct {
	Action     string `json:"action"`
	Controller string `json:"controller"`
	Session    string `json:"session"`
	Params     any    `json:"params"`
}

func (p *payload) GetBody() (io.Reader, error) {
	body, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(body), nil
}

func (p *payload) GetHeader() request.Header {
	return request.Header{"Content-Type": contentType}
}

type cacheParams struct {
	Names []string `json:"names"`
}
