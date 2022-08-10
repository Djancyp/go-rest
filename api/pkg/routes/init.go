package routers

import "github.com/gorilla/mux"

func RouterInit(router *mux.Router) {
	RegisterExampleRoutes(router)
	ImportMiddlewares(router)
}
