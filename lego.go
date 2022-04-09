package main

import (
    "log"

    "github.com/didadadida93/lego/pkg/request"
)

func main() {
    params := request.UrlParams{
        "param1": "value1",
        "param2": "value2",
    }
    res, err := request.Get("https://httpbin.org/get", params)
    if err != nil {
        log.Fatal(err)
        return
    }

    log.Println(res)

    res1, err := request.Get("https://httpbin.org/get", nil)
    if err != nil {
        log.Fatal(err)
        return
    }

    log.Println(res1)
}
