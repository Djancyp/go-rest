package main

import (
	"fmt"
	"github.com/Djancyp/go-rest/pkg"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		pkg.PrintMe()
	})
	http.ListenAndServe(":8080", nil)
}
