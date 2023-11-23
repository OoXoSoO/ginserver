package nethttp

import (
	"ginserver/server"
	"net/http"
)

type Router struct {
	routes     map[string]map[string]http.HandlerFunc
	middleware []server.Middleware
}

// ensure server.Route implementation
var _ server.Router = (*Router)(nil)

func NewRouter(middlewares ...server.Middleware) *Router {
	return &Router{
		routes:     make(map[string]map[string]http.HandlerFunc),
		middleware: middlewares,
	}
}

func (gr *Router) AddRoute(method, path string, handlerFunc http.HandlerFunc, middlewares ...server.Middleware) {
	if gr.routes[method] == nil {
		gr.routes[method] = make(map[string]http.HandlerFunc)
	}
	handlerFunc = applyMiddleware(handlerFunc, middlewares...)
	gr.routes[method][path] = handlerFunc
}

func (gr *Router) Start(port string) {
	mainHfn := func(w http.ResponseWriter, r *http.Request) {
		handlerFunc, ok := gr.routes[r.Method][r.URL.Path]
		if ok {
			handlerFunc(w, r)
		} else {
			http.NotFound(w, r)
		}
	}
	mainHfn = applyMiddleware(mainHfn, gr.middleware...)
	http.HandleFunc("/", mainHfn)
	http.ListenAndServe(":"+port, nil)
}

func applyMiddleware(handlerFunc http.HandlerFunc, middlewares ...server.Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		handlerFunc = middleware(handlerFunc)
	}
	return handlerFunc
}
