package response

import (
	"io"
	"net/http"
)

type Response struct {
	Header http.Header
	Body   string
}

func NewResponse(res *http.Response) (r *Response, err error) {
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	r = &Response{res.Header, string(body)}
	return
}
