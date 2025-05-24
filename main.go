package main

import (
	"blog-api/database"
	"blog-api/handler"
	"blog-api/repository"
	"blog-api/routes"
	"blog-api/services"
	"fmt"
	"net/http"
)

func main() {
	// connect to database
	database.ConnectDB() 

	// initiliase the repo
	userRepo := &repository.UserRepo{}

	// initialise the service
	userService := &services.UserService{Repo: userRepo}

	// initialise the handler
	userHandler := &handler.UserHandler{Service: userService}

	// define routes
	router := routes.SetupRouter(userHandler)

	// start the server
	fmt.Println("server is running on localhost:8080...")
	http.ListenAndServe(":8080", router)
}