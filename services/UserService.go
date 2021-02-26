package services

import (
	"net/http"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (h *UserService) LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write([]byte(`{"message" : "This function  is not implemented Yet"}`))
}
func (h *UserService) GetUserByID(id uint) (string, error) {
	return "RANAK", nil
}
