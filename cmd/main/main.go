package main

import (
	"fmt"
	"github.com/Djancyp/go-rest/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	// register routers
	exampleRoutes.RegisterExampleRoutes(r)
	//end of register routers

	http.Handle("/", r)
	fmt.Println("Server is running, http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
