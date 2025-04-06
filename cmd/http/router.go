package http

type HandlerFunc func(string, string)

type Router struct {
	handlers map[string]HandlerFunc
}

func (r *Router) Handler(pattern string) HandlerFunc {
	return r.handlers[pattern]
}

func (r *Router) ServeHTTP(res, req string) {
	r.Handler(req)(res, req)
}

func NewRouter() *Router {
	return &Router{
		handlers: make(map[string]HandlerFunc),
	}
}

var DefaultRouter *Router = NewRouter()

func HandleFunc(pattern string, handlerFunc HandlerFunc) {
	DefaultRouter.handlers[pattern] = handlerFunc
}
