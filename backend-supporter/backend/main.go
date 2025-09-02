package main

import (
	app_handler "backend-supporter/backend/api/app/handler"
	"backend-supporter/backend/config"
	"backend-supporter/backend/middleware"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config)

	r := mux.NewRouter()
	app_handler.Register(r)
	middleware.Register(r)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/dist")))

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
