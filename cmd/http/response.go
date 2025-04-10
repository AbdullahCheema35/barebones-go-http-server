package http

import "fmt"

type Response struct {
	StatusCode    int
	Status        string
	Proto         string
	ProtoMajor    int
	ProtoMinor    int
	ContentLength int64
	Headers       Header
	Body          string
	Request       *Request
}

func (r *Response) setResponseStatus(code int) {
	r.StatusCode = code
	r.Status = getStatusFromCode(code)
}

func (r *Response) getEscapedStatusHeader() string {
	return fmt.Sprintf("%s %d %s\r\n", r.Proto, r.StatusCode, r.Status)
}

func (r *Response) getOtherHeaders() []string {
	return r.Headers.getAllHeaders()
}

func newResponse(request *Request) *Response {
	return &Response{
		Request: request,
		Proto:   supportedHttpProtocol,
		Headers: newHeaderWithTypicalValues(),
	}
}

func getStatusFromCode(code int) string {
	switch code {
	case 200:
		return "OK"
	case 404:
		return "Not Found"
	case 400:
		return "Bad Request"
	default:
		return "Unknown Code"
	}
}
