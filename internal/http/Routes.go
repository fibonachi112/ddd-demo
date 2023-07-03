package http

import (
	"repo/internal/http/handlers"
	_ "net/http/pprof"
)

func (a *HttpApp) makeRoutes() {

	v1 := a.FiberApp.Group("/v1")

	//-------- users routes
	users := v1.Group("/users")
	users.Post("/create", handlers.UserCreateOne)
	users.Post("/update", handlers.UserCreateOne)
	users.Post("/get/:user_id", handlers.UserGetOne)

	//-------- customerRoutes
	customers := v1.Group("/customers")
	customers.Post("/create",handlers.CustomerCreateOne)
	customers.Post("/get/:customer_id",handlers.CustomerCreateOne)
	customers.Post("/update",handlers.CustomerCreateOne)
}
