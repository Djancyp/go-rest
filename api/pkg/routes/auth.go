package routers

import (
	"github.com/Djancyp/go-rest/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterAuthRouters = func(router *mux.Router) {

	router.HandleFunc("/login", controllers.LoginAuth).Methods("POST")
	router.HandleFunc("/register", controllers.AuthRegister).Methods("POST")
	router.HandleFunc("/reset-password", controllers.Auth(controllers.ChangePassword)).Methods("POST")
	router.HandleFunc("/refresh", controllers.Auth(controllers.Refresh)).Methods("GET")
	router.HandleFunc("/forgot-password", controllers.PassworRecovery).Methods("POST")
	//TODO add route for add, delete ,update types
	auth_role := []controllers.AuthRole{
		{ID: 1, Role: "superuser"},
	}
	router.HandleFunc("/auth-role", controllers.AuthRoles(controllers.AddRole, auth_role)).Methods("POST")
}
