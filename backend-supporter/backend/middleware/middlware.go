package middleware

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	r.Use(loggingMiddleware)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
