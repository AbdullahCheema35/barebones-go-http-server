package http

type Header map[string]string

func (h Header) Set(key, val string) {
	h[key] = val
}
