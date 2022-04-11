package request

import (
	"io"
	"net/http"
	"net/url"
)

type UrlParams map[string]string

type RequestConfig struct {
	Url            string
	Params         UrlParams
	Body           IPayload
	FollowRedirect bool
	Method         string
}

func (rc *RequestConfig) GetBody() (io.Reader, error) {
	if rc.Method == "GET" {
		return nil, nil
	}
	return rc.Body.GetBody()
}

func (rc *RequestConfig) GetUrl() (string, error) {
	if rc.Params == nil {
		return rc.Url, nil
	}

	u, err := url.Parse(rc.Url)
	if err != nil {
		return "", err
	}

	q := u.Query()
	for k, v := range rc.Params {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}

func (rc *RequestConfig) GetHeaders() RequestHeaders {
	if rc.Method == "GET" {
		return nil
	}
	return rc.Body.GetHeaders()
}

func (rc *RequestConfig) PrepareRequest() (*http.Request, error) {
	var reqBody io.Reader = nil

	url, err := rc.GetUrl()
	if err != nil {
		return nil, err
	}

	reqBody, err = rc.GetBody()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(rc.Method, url, reqBody)
	if err != nil {
		return nil, err
	}

	setRequestHeader(req, rc.GetHeaders())

	return req, nil
}

func setRequestHeader(req *http.Request, headers RequestHeaders) {
	// set header user-agent manually
	req.Header.Set("User-Agent", userAgent)

	if headers == nil {
		return
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}
}
