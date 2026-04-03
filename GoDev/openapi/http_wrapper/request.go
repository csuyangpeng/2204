package http_wrapper

import (
	"net/http"
	"net/url"
)

type Request struct {
	Params map[string]string
	Header http.Header
	Query  url.Values
	Body   interface{}
}

func NewRequest(req *http.Request, body interface{}) (ret *Request) {
	ret = &Request{}

	//todo parse the reqest and filled the params, header,query etc.
	ret.Header = req.Header
	ret.Body = body
	ret.Params = make(map[string]string)
	ret.Query = url.Values(make(map[string][]string))
	return
}
