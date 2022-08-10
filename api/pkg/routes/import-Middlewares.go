package routers

import (
	"github.com/Djancyp/go-rest/pkg/middleware"
	"github.com/gorilla/mux"
)

var ImportMiddlewares = func(router *mux.Router) {
	//midlewares
	router.Use(middlewares.ReguestLogger)
	router.Use(middlewares.Cors)
}
