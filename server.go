package main

import (
	"net/http"
	"github.com/go-chi/chi"
	"GolangWeb/routers"

	
)

func main() {

	port := ":8090"
	r := chi.NewRouter()
	routers.CustomPrint();
	r.Get("/", routers.Delhandle)
	r.Post("/", routers.Posthandle)
	r.Put("/", routers.Puthandle)
	r.Delete("/", routers.Delhandle)
	http.ListenAndServe(port, r)
	
}
