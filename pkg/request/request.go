package request

import (
    "io"
    "net/http"
)

type UrlParams map[string]string
type RequestHeaders map[string]string

var client *http.Client = &http.Client{}

func Get(u string, p UrlParams) (r string, err error) {
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

    defer res.Body.Close()
    body, err := io.ReadAll(res.Body)
    if err != nil {
        return
    }

    r = string(body)
    return
}

func PostForm(u string, p UrlParams, b io.Reader) (r string, err error) {
    req, err := http.NewRequest("POST", u, b)
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

    defer res.Body.Close()
    body, err := io.ReadAll(res.Body)
    if err != nil {
        return
    }

    r = string(body)
    return
}

func PostJson(u string, p UrlParams, b IPayload) (r string, err error) {
    payload, err := b.GetJsonPayload()
    if err != nil {
        return
    }

    req, err := http.NewRequest("POST", u, payload)
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

    defer res.Body.Close()
    body, err := io.ReadAll(res.Body)
    if err != nil {
        return
    }

    r = string(body)
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
    req.Header.Set("User-Agent", "lego 1.0.0")

    if headers == nil {
        return
    }

    for k, v := range headers {
        req.Header.Set(k, v)
    }
}
