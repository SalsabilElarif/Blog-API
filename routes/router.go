package routes

import (
	"blog-api/handler"
	"github.com/gorilla/mux"
)

func SetupRouter(userHandler *handler.UserHandler, postHandler *handler.PostHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")
	r.HandleFunc("/login", userHandler.Login).Methods("POST")

	return r
}
