package main

import (
	"blog-api/config"
	"blog-api/database"
	"blog-api/handler"
	"blog-api/repository"
	"blog-api/routes"
	"blog-api/services"
	"fmt"
	"net/http"
)

func main() {
	// load up variables
	config.LoadEnv()

	// connect to database
	database.ConnectDB() 

	// initiliase the repo
	userRepo := &repository.UserRepo{}
	postRepo := &repository.PostRepo{}

	// initialise the service
	userService := &services.UserService{Repo: userRepo}
	postService := &services.PostService{Repo: postRepo}


	// initialise the handler
	userHandler := &handler.UserHandler{Service: userService}
	postHandler := &handler.PostHandler{Service: postService}
	// define routes
	router := routes.SetupRouter(userHandler, postHandler)

	// start the server
	fmt.Println("server is running on localhost:8080...")
	http.ListenAndServe(":8080", router)
}