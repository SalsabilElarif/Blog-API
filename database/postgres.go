package database

import (
	"blog-api/models"
	"fmt"
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// connection string
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", 
		os.Getenv("DB_HOST"), 
		os.Getenv("DB_USER"), 
		os.Getenv("DB_PASSWORD"), 
		os.Getenv("DB_NAME"), 
		os.Getenv("DB_PORT"),
	)

	// log.Println("env...", connStr)

	var err error
	// open database connection
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// migrate models to form tables in the database
	err = DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	if err != nil {
		log.Fatalf("failed to migrate models: %v", err)
	}
	fmt.Println("connected to database successfully")
} 