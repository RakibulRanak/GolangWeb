package handlers

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi"	
	"GolangWeb/services"
	"strconv"
	"encoding/json"
)
type UserHandler struct {
}
type myresponse struct{
	Name string `json:"name"`
}
func NewUserHandler() *UserHandler {
	
	return &UserHandler{}
}

func (h *UserHandler) Handle(router chi.Router) {
	
	router.Route("/{id}", func(router chi.Router) {
		router.Get("/", h.getUserByID)
		//router.Delete("/", h.deleteUser)
	})
	router.Get("/all", h.getUser)  
	router.Post("/", h.createUser) 
	router.Put("/", h.updateUser)
	router.Delete("/", h.deleteUser)
	userService := services.NewUserService()
	router.Post("/login",userService.LoginUser)
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

func (h *UserHandler) getUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	fmt.Println(err)
	Id:=uint(id)
	userService := services.NewUserService()
	d, e := userService.GetUserByID(Id)
	if e != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte(`{"name" : "Not found"}`))
		fmt.Println(d)
	}else{
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		resp1:=myresponse{Name:d}
		byteArray,err := json.Marshal(resp1)
		if err==nil{
		w.Write([]byte(byteArray))
		}
		fmt.Println(d)
	}
}