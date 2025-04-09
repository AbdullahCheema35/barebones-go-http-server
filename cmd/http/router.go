package http

import (
	"strings"
)

type Router struct {
	handlers map[string]Handler
}

func (r *Router) Handler(pattern string) Handler {
	handler, ok := r.handlers[pattern]
	if !ok {
		return notFoundHandler
	}
	return handler
}

func (r *Router) ServeHTTP(res ResponseWriter, req *Request) {
	trimmedEndpoint := strings.TrimSuffix(req.Endpoint, "/")
	if trimmedEndpoint == "" {
		trimmedEndpoint = "/"
	}
	handler := r.Handler(trimmedEndpoint)
	handler.ServeHTTP(res, req)
}

func NewRouter() *Router {
	return &Router{
		handlers: make(map[string]Handler),
	}
}

var DefaultRouter *Router = NewRouter()

func HandleFunc(pattern string, f func(ResponseWriter, *Request)) {
	DefaultRouter.handlers[pattern] = HandlerFunc(f)
}

func Handle(pattern string, h Handler) {
	DefaultRouter.handlers[pattern] = h
}

var notFoundHandler = HandlerFunc(func(resp ResponseWriter, req *Request) {
	resp.WriteHeader(404)
	resp.Write([]byte("HTTP Error: 404\n\nDoes Not Exist"))
})
