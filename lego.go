package main

import "github.com/didadadida93/lego/pkg/login"

func main() {
	login.Login()
	// params := request.UrlParams{
	// "param1": "value1",
	// "param2": "value2",
	// }
	// payload := &request.FormBody{
	// "email": "email@meail.com",
	// "password": "password",
	// }
	// res, err := request.PostForm("https://httpbin.org/post", params, payload, false)
	// if err != nil {
	// log.Fatal(err)
	// return
	// }

	// log.Println(res.Body)

	// res1, err := request.Get("https://httpbin.org/get", params, false)
	// if err != nil {
	// log.Fatal(err)
	// return
	// }

	// log.Println(res1.Body)

	// user := &request.User{"name", 8}
	// res2, err := request.PostJson("https://httpbin.org/post", nil, user, false)
	// if err != nil {
	// log.Fatal(err)
	// return
	// }

	// log.Println(res2.Body)
}
