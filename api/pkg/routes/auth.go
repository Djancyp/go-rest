package routers

import (
	"github.com/Djancyp/go-rest/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterAuthRouters = func(router *mux.Router) {

	router.HandleFunc("/login", controllers.LoginAuth).Methods("POST")
	router.HandleFunc("/register", controllers.AuthRegister).Methods("POST")
	router.HandleFunc("/refresh", controllers.Refresh).Methods("GET")
	router.HandleFunc("/forgot-password", controllers.PassworRecovery).Methods("POST")
}
