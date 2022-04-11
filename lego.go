package main

import (
	// "log"
	// "net/http"

	"github.com/didadadida93/lego/pkg/login"
	// "github.com/didadadida93/lego/pkg/request"
)

func main() {
	login.Login()
	// params := request.UrlParams{
	// "params1": "values1",
	// "params2": "values2",
	// }

	// payload1 := &request.FormBody{
	// "email":    "email@email.com",
	// "password": "password",
	// }

	// rc := request.NewRequestConfig()
	// rc.Set("url", "https://httpbin.org/post")
	// rc.Set("params", params)
	// rc.Set("body", payload1)
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
	// res, err = request.Do(rc)
	// if err != nil {
	// log.Fatal(err)
	// }
	// log.Println(res.Body)
}
