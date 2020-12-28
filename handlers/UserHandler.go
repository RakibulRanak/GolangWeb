package handlers

import (
	"net/http"
	"github.com/go-chi/chi"	
	"GolangWeb/services"
)
type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	
	return &UserHandler{}
}

func (h *UserHandler) Handle(router chi.Router) {
	
	router.Get("/all", h.getUser)  
	router.Post("/", h.createUser) 
	router.Put("/", h.updateUser)
	router.Delete("/", h.deleteUser)
	authService := services.NewAuthService()
	router.Post("/login",authService.LoginUser)
}

func (h *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write([]byte(`{"message" : "data created"}`))
}

func (h *UserHandler) getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(`{"message" : "data fetch"}`))
}

func (h *UserHandler) updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(`{"message" : "data updated"}`))
}

func (h *UserHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(204)
}

