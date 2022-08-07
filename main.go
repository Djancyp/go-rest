package main

import (
	"fmt"
	"github.com/Djancyp/go-rest/pkg"
	"io"
	"net/http"
	"time"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, time.Now().Format("2006-01-02 15:04:05"))
	io.WriteString(w, "Hello world")
	pkg.PrintMe()
}
func main() {
	http.HandleFunc("/", MainHandler)
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
