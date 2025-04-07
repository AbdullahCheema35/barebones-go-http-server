package http

import "strings"

type Request struct {
	Method   string
	Endpoint string
	Proto    string
}

func parseRequest(req string) *Request {
	fistLine := strings.Split(req, "\n")
	parts := strings.Split(fistLine[0], " ")

	return &Request{
		Method:   parts[0],
		Endpoint: parts[1],
		Proto:    parts[2],
	}
}
