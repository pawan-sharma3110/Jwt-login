package router

import (
	"jwt/handler"

	"github.com/gorilla/mux"
)

func MyRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/user/register", handler.RegisterUser).Methods("POST")
	r.HandleFunc("/user/login", handler.Login).Methods("POST")
	return r
}
