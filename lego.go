package main

import (
	// "fmt"
	"log"
	"net/http"
	// "time"

	"github.com/didadadida93/lego/pkg/config"
	"github.com/didadadida93/lego/pkg/request"
)

func main() {
	c, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	gs := c.Authenticate()
  controller, action := "cache", "get"
	// url := fmt.Sprintf("https://%s.kingdoms.com/api/?c=%s&a=%s&t%v",
		// c.Gameworld, controller, action, time.Now().Unix())
	rc := request.NewRequestConfig()
  rc.Set("url", "https://httpbin.org/post")
	rc.Set("params", nil)
	rc.Set("body", &request.TKPayload{
		Action:     action,
		Controller: controller,
		Session:    gs.GameworldSession,
		Params: request.TKParams{
			Names: []string{"Collection:Village:own"},
		},
	})
	rc.Set("header", request.Header{
		"Cookie": gs.GetGameCookie(),
	})
	rc.Set("method", http.MethodPost)
	rc.Set("followRedirect", false)

	res, err := request.Do(rc)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.Body)

	// for testing purpose
	// params := request.UrlParams{
	// "params1": "values1",
	// "params2": "values2",
	// }

	// header := request.Header{
	// "Cookie": "temp cookie from lego",
	// }

	// payload1 := &request.FormBody{
	// "email":    "email@email.com",
	// "password": "password",
	// }

	// rc := request.NewRequestConfig()
	// rc.Set("url", "https://httpbin.org/post")
	// rc.Set("params", params)
	// rc.Set("body", payload1)
	// rc.Set("header", header)
	// rc.Set("followRedirect", false)
	// rc.Set("method", http.MethodPost)

	// res, err := request.Do(rc)
	// if err != nil {
	// log.Fatal(err)
	// }
	// log.Println(res.Body)

	// payload2 := &request.User{
	// Name: "name",
	// Age:  8,
	// }
	// rc.Set("body", payload2)
	// res, err = request.Do(rc)
	// if err != nil {
	// log.Fatal(err)
	// }
	// log.Println(res.Body)

	// rc.Set("url", "https://httpbin.org/get")
	// rc.Set("params", nil)
	// rc.Set("body", nil)
	// rc.Set("method", http.MethodGet)
	// res, err = request.Do(rc)
	// if err != nil {
	// log.Fatal(err)
	// }
	// log.Println(res.Body)

	// rc.Set("params", params)
	// rc.Set("header", nil)
	// res, err = request.Do(rc)
	// if err != nil {
	// log.Fatal(err)
	// }
	// log.Println(res.Body)
}
