package request

import (
	"fmt"
	"io"
	"net/http"

	"github.com/didadadida93/lego/pkg/response"
)

type IPayload interface {
	GetBody() (io.Reader, error)
	GetHeader() Header
}

type UrlParams map[string]string
type Header map[string]string
type Cookie map[string]string

func (c Cookie) String() (s string) {
	for k, v := range c {
		s += fmt.Sprintf("%s=%v; ", k, v)
	}
	return
}

var client *http.Client = &http.Client{CheckRedirect: checkRedirect}

func Do(rc *RequestConfig) (r *response.Response, err error) {
	req, err := rc.PrepareRequest()
	if err != nil {
		return
	}

	setFollowRedirect(rc.FollowRedirect)
	res, err := client.Do(req)
	if err != nil {
		return
	}

	r, err = response.NewResponse(res)
	return
}

func checkRedirect(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}

func setFollowRedirect(b bool) {
	switch b {
	case false:
		client.CheckRedirect = checkRedirect
	default:
		client.CheckRedirect = nil
	}
}
