// in cmd/api/main.go
package main

// @title KEEPER API
// @version 1.0
// @description API for the KEEPER Dealership Management System.
// @host progetto-keeper.fly.dev
// @BasePath /
import (
	"keeper/internal/api"
	"keeper/internal/storage"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using OS environment variables")
	}
	// Get the database connection string from environment variables
	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}
	// Initialize the storage
	store, err := storage.NewPostgresStore(connString)
	if err != nil {
		log.Fatal("failed to connect to the database: ", err)
	}
	// Initialize the validator
	validate := validator.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server := api.NewAPIServer(":"+port, store, validate)
	server.Run()
}
