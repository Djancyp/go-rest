package main

import (
	"fmt"
	"github.com/Djancyp/go-rest/pkg/config"
	"github.com/Djancyp/go-rest/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	port := config.GetConfig("API_PORT")
	api_domain := config.GetConfig("API_DOMAIN")
	r := mux.NewRouter()
	// register routers
	//end of register routers
	routers.RouterInit(r)
	fmt.Println("========================================")
	fmt.Printf("You can access website:, https://%s", api_domain)
	fmt.Println("")
	fmt.Printf("Server is running, http://localhost:%s", port)
	fmt.Println("")
	fmt.Println("========================================")
	fmt.Println("")
	log.Fatal(http.ListenAndServe(":"+port, r))

}
