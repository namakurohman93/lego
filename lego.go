package main

import (
	"log"
	"net/http"

	// "github.com/didadadida93/lego/pkg/login"
	"github.com/didadadida93/lego/pkg/request"
)

func main() {
	// new request with RequestConfig
	params := request.UrlParams{
		"params1": "values1",
		"params2": "values2",
	}
	payload1 := &request.FormBody{
		"email":    "email@email.com",
		"password": "password",
	}
	requestConfig := &request.RequestConfig{
		Url:            "https://httpbin.org/post",
		Params:         params,
		Body:           payload1,
		FollowRedirect: false,
		Method:         http.MethodPost,
	}
	res, err := request.Do(requestConfig)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.Body)

	payload2 := &request.User{
		Name: "name",
		Age:  8,
	}
	requestConfig1 := &request.RequestConfig{
		Url:            "https://httpbin.org/post",
		Params:         params,
		Body:           payload2,
		FollowRedirect: false,
		Method:         http.MethodPost,
	}
	res, err = request.Do(requestConfig1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.Body)

	requestConfig2 := &request.RequestConfig{
		Url:            "https://httpbin.org/get",
		Params:         nil,
		Body:           nil,
		FollowRedirect: false,
		Method:         http.MethodGet,
	}
	res, err = request.Do(requestConfig2)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.Body)

	requestConfig3 := &request.RequestConfig{
		Url:            "https://httpbin.org/get",
		Params:         params,
		Body:           nil,
		FollowRedirect: false,
		Method:         http.MethodGet,
	}
	res, err = request.Do(requestConfig3)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.Body)
}
