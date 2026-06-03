package sbicmn

import (
	"net/http"
)

type Response struct {
	Header http.Header
	Status int
	Body   interface{}
}

func NewResponse(code int, h http.Header, body interface{}) (ret *Response) {
	ret = &Response{}

	// todo check the input parameters
	ret.Header = h
	ret.Status = code
	ret.Body = body
	return ret
}
