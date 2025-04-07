package http

import (
	"fmt"
	"net"
)

type ResponseWriter interface {
	Header() Header
	Write([]byte) (int, error)
	WriteHeader(int)
}

type Response struct {
	Conn       net.Conn
	StatusCode int
	Status     string
	RespHeader Header
	Body       string
}

func (r *Response) Write(data []byte) (int, error) {
	n, err := fmt.Fprintf(r.Conn, "%s", data)
	return n, err
}

func (r *Response) WriteHeader(code int) {
	// Implement
}

func (r *Response) Header() Header {
	return r.RespHeader
}
