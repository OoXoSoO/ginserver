package internal

import (
	"flag"
	"fmt"
	"ginserver/adapters/inmem"
	"ginserver/handler"
	"ginserver/pkg/service"
	"ginserver/server"
	"ginserver/server/gin"
	"ginserver/server/nethttp"
)

func init() {
	flgs = serverflags{}
	framework := ""
	flag.StringVar(&framework, "framework", "", fmt.Sprintf(`determines wich web server framework will be used (possible values "%s","%s")`, FRW_GIN, FRW_NETHTTP))
	flag.Parse()
	if framework != FRW_GIN && framework != FRW_NETHTTP {
		panic("specified framework not valid")
	}
	flgs.Framework = framework
}

var flgs serverflags

type serverflags struct {
	Framework string
}

const (
	FRW_GIN     string = "GIN"
	FRW_NETHTTP string = "NETHTTP"
)

type Infrastrcture struct {
	Router server.Router

	CreateUserHandler handler.UserHandler
	PongHandler       handler.PongHandler
}

func Startup() Infrastrcture {

	inmemRepository := inmem.NewInmemRepository()

	createHandler := handler.UserHandler{
		Service: service.NewUser(inmemRepository),
	}
	pongHandler := handler.PongHandler{
		Service: &service.Pong{},
	}
	return Infrastrcture{
		CreateUserHandler: createHandler,
		PongHandler:       pongHandler,

		Router: buildRouter(),
	}
}

func buildRouter() server.Router {

	switch flgs.Framework {
	case FRW_GIN:
		return gin.NewGinRouter()
	case FRW_NETHTTP:
		return nethttp.NewRouter()
	}
	panic("init framework not valid")
}
