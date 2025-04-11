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
		return r.NotFoundHandler()
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

func (r *Router) HandleFunc(pattern string, f func(ResponseWriter, *Request)) {
	r.Handle(pattern, HandlerFunc(f))
}

func (r *Router) Handle(pattern string, h Handler) {
	r.handlers[pattern] = h
}

func (r *Router) NotFoundHandler() Handler {
	h, ok := r.handlers["404"]
	if !ok {
		return defaultNotFoundHandler
	}
	return h
}

func (r *Router) SetNotFoundHandler(h Handler) {
	r.Handle("404", h)
}

func (r *Router) SetNotFoundHandlerFunc(f func(ResponseWriter, *Request)) {
	r.SetNotFoundHandler(HandlerFunc(f))
}

func NewRouter() *Router {
	return &Router{
		handlers: make(map[string]Handler),
	}
}

var DefaultRouter *Router = NewRouter()

func HandleFunc(pattern string, f func(ResponseWriter, *Request)) {
	DefaultRouter.HandleFunc(pattern, f)
}

func Handle(pattern string, h Handler) {
	DefaultRouter.Handle(pattern, h)
}

var defaultNotFoundHandler = HandlerFunc(func(resp ResponseWriter, req *Request) {
	resp.WriteHeader(404)
	resp.Write([]byte("HTTP Error: 404\n\nDoes Not Exist"))
})
