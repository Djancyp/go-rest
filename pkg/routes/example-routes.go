package exampleRoutes

import (
	"github.com/Djancyp/go-rest/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterExampleRoutes = func(router *mux.Router) {
	router.HandleFunc("/example", controllers.GetAllExamples).Methods("GET")
	// router.HandleFunc("/example", controllers.CreateExample).Methods("POST")
}
