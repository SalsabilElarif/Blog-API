package main

import (
	"blog-api/database"
	"fmt"
	"net/http"

)

func main() {
	// Connect to database
	database.ConnectDB() 

	// Define routes

	// Start the server
	fmt.Println("server is running on localhos:8080...")
	http.ListenAndServe(":8080", nil)
}