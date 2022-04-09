package request

import (
    "io"
    "net/http"
)

type UrlParams map[string]string

var client *http.Client = &http.Client{}

func Get(url string, params UrlParams) (result string, err error) {
    req, err := http.NewRequest("GET", url, nil)
    if params != nil {
        addUrlParams(req, params)
    }

    res, err := client.Do(req)
    if err != nil {
        return
    }

    defer res.Body.Close()
    body, err := io.ReadAll(res.Body)
    if err != nil {
        return
    }

    result = string(body)
    return
}

func addUrlParams(req *http.Request, params UrlParams) {
    q := req.URL.Query()
    for k, v := range params {
        q.Add(k, v)
    }
    req.URL.RawQuery = q.Encode()
}
