package main

import (
	"net/http"
	"github.com/go-chi/chi"
	"GolangWeb/handlers"

	
)

func main() {

	port := ":8090"
	r := chi.NewRouter()
	userHandler := handlers.NewUserHandler()
	r.Route("/users", userHandler.Handle)
	http.ListenAndServe(port, r)
	
}
