package routers

import (
	"github.com/Djancyp/go-rest/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterExampleRoutes = func(router *mux.Router) {
	//routers
	//Example fo auth middleware with JWT
	router.HandleFunc("/", controllers.Auth(controllers.GetAllExamples)).Methods("GET")
	router.HandleFunc("/example", controllers.GetAllExamples).Methods("GET")
	router.HandleFunc("/example/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/example", controllers.CreateExample).Methods("POST")
	router.HandleFunc("/example/{id}", controllers.DeleteExample).Methods("DELETE")
	router.HandleFunc("/example/{id}", controllers.UpdateExampleById).Methods("PUT")
}
