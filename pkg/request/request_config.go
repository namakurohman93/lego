package request

import (
	"io"
	"net/http"
	"net/url"
)

type UrlParams map[string]string
type RequestHeader map[string]string

type RequestConfig struct {
	url            string
	params         UrlParams
	body           IPayload
	header         RequestHeader
	FollowRedirect bool
	method         string
}

func (rc *RequestConfig) GetBody() (io.Reader, error) {
	if rc.method == "GET" {
		return nil, nil
	}

	if rc.body == nil {
		return nil, nil
	}

	return rc.body.GetBody()
}

func (rc *RequestConfig) GetUrl() (string, error) {
	if rc.params == nil {
		return rc.url, nil
	}

	u, err := url.Parse(rc.url)
	if err != nil {
		return "", err
	}

	q := u.Query()
	for k, v := range rc.params {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}

func (rc *RequestConfig) GetHeader() RequestHeader {
	if rc.method == "GET" {
		return nil
	}

	if rc.body == nil {
		return nil
	}

	h := rc.body.GetHeader()

	if rc.header != nil {
		// merge header
		mergeHeader(&h, &rc.header)
	}

	return h
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

	req, err := http.NewRequest(rc.method, url, reqBody)
	if err != nil {
		return nil, err
	}

	setRequestHeader(req, rc.GetHeader())

	return req, nil
}

func (rc *RequestConfig) Set(key string, value any) (ok bool) {
	ok = true
	switch key {
	case "url":
		if value == nil {
			rc.url = ""
		} else {
			rc.url = value.(string)
		}
	case "params":
		if value == nil {
			rc.params = nil
		} else {
			rc.params = value.(UrlParams)
		}
	case "body":
		if value == nil {
			rc.body = nil
		} else {
			rc.body = value.(IPayload)
		}
	case "followRedirect":
		if value == nil {
			rc.FollowRedirect = false
		} else {
			rc.FollowRedirect = value.(bool)
		}
	case "method":
		if value == nil {
			rc.method = "GET"
		} else {
			rc.method = value.(string)
		}
	case "header":
		if value == nil {
			rc.header = nil
		} else {
			rc.header = value.(RequestHeader)
		}
	default:
		ok = false
	}
	return
}

func NewRequestConfig() *RequestConfig {
	return &RequestConfig{}
}

func setRequestHeader(req *http.Request, headers RequestHeader) {
	// set header user-agent manually
	req.Header.Set("User-Agent", userAgent)

	if headers == nil {
		return
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}
}

func mergeHeader(d, s *RequestHeader) {
	for k, v := range *s {
		(*d)[k] = v
	}
}
