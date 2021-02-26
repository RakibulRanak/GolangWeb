package main

import (
	"GolangWeb/db"
	"GolangWeb/handlers"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

func main() {

	port := ":8090"
	r := chi.NewRouter()
	userHandler := handlers.NewUserHandler()
	r.Route("/users", userHandler.Handle)
	http.ListenAndServe(port, r)
	dbase := db.GetDB()
	fmt.Print(dbase)
}
