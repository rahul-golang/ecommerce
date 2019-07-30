package endpoints

import (
	"github.com/gorilla/mux"
	"github.com/rahul-golang/ecommerce/users/pkg/httphandler"
)

func NewUserEndpoints(router *mux.Router, restHandler *httphandler.UserHttpHandlers) {
	router.HandleFunc("/users/{id}", restHandler.GetUser).Methods("GET")
	router.HandleFunc("/users", restHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users", restHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/users", restHandler.DeleteUser).Methods("Delete")
	router.HandleFunc("/users", restHandler.GetAllUser).Methods("GET")

}
