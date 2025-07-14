package routes

import (
	"blog-api/handler"
	"blog-api/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter(userHandler *handler.UserHandler, postHandler *handler.PostHandler) *mux.Router {
	r := mux.NewRouter()

	//public routes
	r.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")
	r.HandleFunc("/login", userHandler.Login).Methods("POST")

	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	//Editor-only routes
	protected.Handle("/posts", middleware.EditorOnlyMiddleware(http.HandlerFunc(postHandler.CreatePost))).Methods("POST")


	//Authenticated user routes
	protected.HandleFunc("/post", postHandler.GetPostByID).Methods("GET")
	
	return r
}

