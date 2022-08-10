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
	//end of register routers
	routers.RouterInit(r)
	fmt.Println("========================================")
	fmt.Println("Server is running, http://localhost:8080")
	fmt.Println("========================================")
	fmt.Println("")
	log.Fatal(http.ListenAndServe(":8080", r))

}
