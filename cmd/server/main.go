package main

import "ginserver/internal"

func main() {
	inf := internal.Startup()

	inf.Router.AddRoute("POST", "/users", inf.CreateUserHandler.CreateUser)
	inf.Router.AddRoute("GET", "/users", inf.CreateUserHandler.GetAll)
	
	inf.Router.AddRoute("GET", "/ping", inf.PongHandler.Pong)

	inf.Router.Start("8080")
}
