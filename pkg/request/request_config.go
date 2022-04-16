package request

import (
	"io"
	"net/http"
	"net/url"
)

type RequestConfig struct {
	url            string
	params         UrlParams
	body           IPayload
	header         Header
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

func (rc *RequestConfig) GetHeader() Header {
	h := make(Header)

	if rc.body != nil {
		t := rc.body.GetHeader()
		mergeHeader(&h, &t)
	}

	if rc.header != nil {
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
	switch v := value.(type) {
	case UrlParams:
		rc.params = v
	case IPayload:
		rc.body = v
	case Header:
		rc.header = v
	case bool:
		rc.FollowRedirect = v
	case string:
		if key == "url" {
			rc.url = v
		} else if key == "method" {
			rc.method = v
		} else {
			ok = false
		}
	default:
		ok = false
	}
	return
}

func NewRequestConfig() *RequestConfig {
	return &RequestConfig{}
}

func setRequestHeader(req *http.Request, headers Header) {
	// set header user-agent manually
	req.Header.Set("User-Agent", userAgent)

	if headers == nil {
		return
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}
}

func mergeHeader(d, s *Header) {
	for k, v := range *s {
		(*d)[k] = v
	}
}
