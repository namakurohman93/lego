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
    payload := request.FormBody{
        "email": "email@meail.com",
        "password": "password",
    }
    formBody := request.GenerateFormBody(payload)
    res, err := request.PostForm("https://httpbin.org/post", params, formBody)
    if err != nil {
        log.Fatal(err)
        return
    }

    log.Println(res)

    res1, err := request.Get("https://httpbin.org/get", params)
    if err != nil {
        log.Fatal(err)
        return
    }

    log.Println(res1)

    user := &request.User{"name", 8}
    res2, err := request.PostJson("https://httpbin.org/post", nil, user)
    if err != nil {
        log.Fatal(err)
        return
    }

    log.Println(res2)
}
