package http

type Router struct {
	handlers map[string]Handler
}

func (r *Router) Handler(pattern string) Handler {
	return r.handlers[pattern]
}

func (r *Router) ServeHTTP(res ResponseWriter, req *Request) {
	handler := r.Handler("/")
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
