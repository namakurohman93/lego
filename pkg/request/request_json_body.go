package request

import (
    "bytes"
    "encoding/json"
    "io"
)

type IPayload interface {
    GetJsonPayload() (io.Reader, error)
}

type User struct {
    Name string `json:"name"`
    Age int `json:"age"`
}

func (u *User) GetJsonPayload() (io.Reader, error) {
    body, err := json.Marshal(u)
    if err != nil {
        return nil, err
    }
    return bytes.NewBuffer(body), nil
}
