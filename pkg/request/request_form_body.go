package request

import (
    "io"
    "net/url"
    "strings"
)

type FormBody map[string]string

func GenerateFormBody(p FormBody) io.Reader {
    d := url.Values{}
    for k, v := range p {
        d.Set(k, v)
    }
    return strings.NewReader(d.Encode())
}
