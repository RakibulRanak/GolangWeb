package services

import (
	"net/http"
)
type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (h *AuthService) LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write([]byte(`{"message" : "This function  is not implemented Yet"}`))
}
