package request

import (
    "io"
    "net/url"
    "strings"
)

type FormBody map[string]string

func (f *FormBody) GetBody() (io.Reader, error) {
    d := url.Values{}
    for k, v := range *f {
        d.Set(k, v)
    }
    return strings.NewReader(d.Encode()), nil
}
