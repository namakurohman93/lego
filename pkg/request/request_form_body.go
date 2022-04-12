package request

import (
	"io"
	"net/url"
	"strings"
)

type FormBody map[string]string

var formContentType string = "application/x-www-form-urlencoded"

func (f *FormBody) GetBody() (io.Reader, error) {
	d := url.Values{}
	for k, v := range *f {
		d.Set(k, v)
	}
	return strings.NewReader(d.Encode()), nil
}

func (f *FormBody) GetHeader() RequestHeader {
	return RequestHeader{"Content-Type": formContentType}
}
