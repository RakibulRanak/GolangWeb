package controllers

import (
	"net/http"
)
type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (h *AuthController) LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write([]byte(`{"message" : "This function  is not implemented Yet"}`))
}
