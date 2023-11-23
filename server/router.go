package server

import "net/http"

type Router interface {
	AddRoute(method, path string, handlerFunc http.HandlerFunc, middlewares ...Middleware)
	Start(port string)
}


