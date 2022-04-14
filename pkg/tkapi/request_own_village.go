package tkapi

import (
	"fmt"
	"net/http"
	"time"

	"github.com/didadadida93/lego/pkg/request"
	"github.com/didadadida93/lego/pkg/response"
)

func RequestOwnVillage(gw, gs, cookie string) (*response.Response, error) {
	controller, action := "cache", "get"
	url := fmt.Sprintf(gameworldUrl, gw, controller,
		action, time.Now().Unix())
	rc := request.NewRequestConfig()
	rc.Set("url", url)
	rc.Set("params", nil)
	rc.Set("body", &request.TKPayload{
		Action:     action,
		Controller: controller,
		Session:    gs,
		Params: request.TKParams{
			Names: []string{"Collection:Village:own"},
		},
	})
	rc.Set("header", request.Header{
		"Cookie": cookie,
	})
	rc.Set("method", http.MethodPost)
	rc.Set("followRedirect", false)

	res, err := request.Do(rc)
	return res, err
}
