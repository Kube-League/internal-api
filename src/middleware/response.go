package middleware

import "encoding/json"

type Response struct {
	Code uint `json:"code"`
	Data any  `json:"data"`
}

func (r *Response) Json() ([]byte, error) {
	return json.Marshal(r)
}
