// in cmd/api/main.go
package main
// @title KEEPER API
// @version 1.0
// @description API for the KEEPER Dealership Management System.
// @host localhost:8080
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

	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	store, err := storage.NewPostgresStore(connString)
	if err != nil {
		log.Fatal("failed to connect to the database: ", err)
	}

	validate := validator.New()
	server := api.NewAPIServer(":8080", store, validate)
	server.Run()
}
