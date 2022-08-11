package routers

import "github.com/gorilla/mux"

func RouterInit(router *mux.Router) {
	ImportMiddlewares(router)
	RegisterExampleRoutes(router)
	RegisterAuthRouters(router)
}
