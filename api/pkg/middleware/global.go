package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func ReguestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("\033[1;36m%s %s\033[0m", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		end := time.Since(start) / (100 * time.Microsecond) * 100 * time.Microsecond
		fmt.Printf("%12v\n", strings.Replace(end.String(), "µ", "u", 1))
	})
}

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		next.ServeHTTP(w, r)
	})
}
