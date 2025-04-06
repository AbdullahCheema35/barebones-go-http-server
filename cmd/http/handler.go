package http

type Handler interface {
	// ServeHTTP(ResponseWriter, *Request)
	ServeHTTP(string, string)
}

// type HandlerFunc func(ResponseWriter, *Request)
// type HandlerFunc func(string, string)

// func (h HandlerFunc) ServeHTTP(res, req string) {

// }
