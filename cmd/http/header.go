package http

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
