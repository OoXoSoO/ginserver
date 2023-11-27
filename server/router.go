package server

import "net/http"

type Router interface {
	// add a new API entry point
	AddRoute(method, path string, handlerFunc http.HandlerFunc, middlewares ...Middleware)

	// Start the server 
	Start(port string)
}


