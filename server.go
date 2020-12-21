package main

import (
	"net/http"
	"github.com/go-chi/chi"
)


func Gethandle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Data fetch"}`))
}
func Posthandle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Data Create"}`))
	
}
func Puthandle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "Data update"}`))
}
func Delhandle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Data delete"}`))
		
}

func main() {

	port := ":8090"
	r := chi.NewRouter()
	r.Get("/", Gethandle)
	r.Post("/", Posthandle)
	r.Put("/", Puthandle)
	r.Delete("/", Delhandle)
	http.ListenAndServe(port, r)
}
