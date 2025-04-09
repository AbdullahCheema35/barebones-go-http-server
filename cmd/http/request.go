package http

import "strings"

type Request struct {
	Method   string
	Endpoint string
	Proto    string
}

func parseRequest(req string) *Request {
	firstLine := strings.Split(req, "\r\n")[0]
	parts := strings.Split(firstLine, " ")

	return &Request{
		Method:   parts[0],
		Endpoint: parts[1],
		Proto:    parts[2],
	}
}
