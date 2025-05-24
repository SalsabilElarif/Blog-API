package routes

import (
	"blog-api/handler"
	"github.com/gorilla/mux"
)

func SetupRouter(userHandler *handler.UserHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")

	return r
}