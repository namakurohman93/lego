package main

import (
	// "fmt"
	"log"
	// "net/http"
	// "time"
	// "strings"

	"github.com/didadadida93/lego/pkg/config"
	"github.com/didadadida93/lego/pkg/tkapi"
	// "github.com/didadadida93/lego/pkg/request"
)

func main() {
	// t1, err := time.Parse(time.RFC1123, "Sun, 24 Apr 2022 08:04:39 GMT")
	// t1, err := time.Parse(time.RFC1123, strings.ReplaceAll("Sun, 24-Apr-2022 08:04:39 GMT", "-", " "))
	// if err != nil {
	// log.Fatal(err)
	// }
	// st := t1.Format(time.RFC3339)
	// t2, err := time.Parse(time.RFC3339, st)
	// if err != nil {
	// log.Fatal(err)
	// }
	// yesterday := t2.AddDate(0, 0, -1)
	// n := time.Now()
	// log.Println(yesterday.Before(n))
	c, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	gs, err := tkapi.Authenticate(&c)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(gs)
	// controller, action := "cache", "get"
	// url := fmt.Sprintf("https://%s.kingdoms.com/api/?c=%s&a=%s&t%v",
	// c.Gameworld, controller, action, time.Now().Unix())
	// rc := request.NewRequestConfig()
	// rc.Set("url", "https://httpbin.org/post")
	// rc.Set("params", nil)
	// rc.Set("body", &request.TKPayload{
	// Action:     action,
	// Controller: controller,
	// Session:    gs.GameworldSession,
	// Params: request.TKParams{
	// Names: []string{"Collection:Village:own"},
	// },
	// })
	// rc.Set("header", request.Header{
	// "Cookie": gs.GetGameCookie(),
	// })
	// rc.Set("method", http.MethodPost)
	// rc.Set("followRedirect", false)

	// res, err := request.Do(rc)
	// if err != nil {
	// log.Fatal(err)
	// }
	// log.Println(res.Body)

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
