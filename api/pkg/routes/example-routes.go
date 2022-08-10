package exampleRoutes

import (
	"github.com/Djancyp/go-rest/pkg/controllers"
	"github.com/Djancyp/go-rest/pkg/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

var RegisterExampleRoutes = func(router *mux.Router) {
	//midlewares
	router.Use(middlewares.ReguestLogger)
	router.Use(middlewares.Cors)
	//routers
	router.HandleFunc("/", controllers.GetAllExamples).Methods("GET")
	router.HandleFunc("/example", controllers.GetAllExamples).Methods("GET")
	router.HandleFunc("/example/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/example", controllers.CreateExample).Methods("POST")
	router.HandleFunc("/example/{id}", controllers.DeleteExample).Methods("DELETE")
	router.HandleFunc("/example/{id}", controllers.UpdateExampleById).Methods("PUT")
}

func Auth(HandlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		HandlerFunc(w, r)
	}
}
