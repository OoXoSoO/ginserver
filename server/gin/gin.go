package gin

import (
	"ginserver/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	engine     *gin.Engine
	middleware []server.Middleware
}

// ensure server.Route implementation
var _ server.Router = (*GinRouter)(nil)

func NewGinRouter(middlewares ...server.Middleware) *GinRouter {
	engine := gin.Default()

	return &GinRouter{
		engine:     engine,
		middleware: middlewares,
	}
}

func (r *GinRouter) AddRoute(method, path string, handlerFunc http.HandlerFunc, middlewares ...server.Middleware) {

	// Apply the general middlewares first
	handlerFunc = applyMiddleware(handlerFunc, r.middleware...)

	// Apply the specific route middlewares
	handlerFunc = applyMiddleware(handlerFunc, middlewares...)

	switch method {
	case "GET":
		r.engine.GET(path, gin.WrapH(handlerFunc))
	case "POST":
		r.engine.POST(path, gin.WrapH(handlerFunc))
	case "PUT":
		r.engine.PUT(path, gin.WrapH(handlerFunc))
	case "DELETE":
		r.engine.DELETE(path, gin.WrapH(handlerFunc))
	default:
		panic("invalid http method")
	}
}

func (r *GinRouter) Start(port string) {
	err := r.engine.Run(":" + port)
	if err != nil {
		panic(err)
	}
}

func applyMiddleware(handlerFunc http.HandlerFunc, middlewares ...server.Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		handlerFunc = middleware(handlerFunc)
	}
	return handlerFunc
}
