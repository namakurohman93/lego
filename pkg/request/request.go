package request

import (
    "io"
    "net/http"

    "github.com/didadadida93/lego/pkg/response"
)

type UrlParams map[string]string
type RequestHeaders map[string]string

type IPayload interface {
    GetBody() (io.Reader, error)
}

var client *http.Client = &http.Client{CheckRedirect: checkRedirect}

func Get(u string, p UrlParams) (r *response.Response, err error) {
    req, err := http.NewRequest("GET", u, nil)
    if err != nil {
        return
    }

    addUrlParams(req, p)

    setRequestHeader(req, nil)

    res, err := client.Do(req)
    if err != nil {
        return
    }

    r, err = response.NewResponse(res)
    return
}

func PostForm(u string, p UrlParams, b IPayload) (r *response.Response, err error) {
    reqBody, err := b.GetBody()
    if err != nil {
        return
    }

    req, err := http.NewRequest("POST", u, reqBody)
    if err != nil {
        return
    }

    addUrlParams(req, p)

    headers := RequestHeaders{
        "Content-Type": "application/x-www-form-urlencoded",
    }
    setRequestHeader(req, headers)

    res, err := client.Do(req)
    if err != nil {
        return
    }

    r, err = response.NewResponse(res)
    return
}

func PostJson(u string, p UrlParams, b IPayload) (r *response.Response, err error) {
    reqBody, err := b.GetBody()
    if err != nil {
        return
    }

    req, err := http.NewRequest("POST", u, reqBody)
    if err != nil {
        return
    }

    addUrlParams(req, p)

    headers := RequestHeaders{
        "Content-Type": "application/json",
    }
    setRequestHeader(req, headers)


    res, err := client.Do(req)
    if err != nil {
        return
    }

    r, err = response.NewResponse(res)
    return
}

func addUrlParams(req *http.Request, params UrlParams) {
    if params == nil {
        return
    }

    q := req.URL.Query()
    for k, v := range params {
        q.Add(k, v)
    }
    req.URL.RawQuery = q.Encode()
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

func checkRedirect(req *http.Request, via []*http.Request) error {
    return http.ErrUseLastResponse
}
