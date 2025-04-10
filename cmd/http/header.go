package http

import "time"

type Header map[string]string

func (h Header) Set(key, val string) {
	h[key] = val
}

func (h Header) getAllHeaders() []string {
	headers := make([]string, 0, len(h))
	for key, val := range h {
		headers = append(headers, key+": "+val)
	}
	return headers
}

func newHeaderWithTypicalValues() Header {
	headers := make(Header)
	headers.Set("Content-Type", "application/json")
	headers.Set("Content-Length", "0")
	headers.Set("Server", "Go HTTP Server From Scratch")
	headers.Set("Date", time.Now().Format(time.RFC1123))
	headers.Set("Connection", "close")
	return headers
}
