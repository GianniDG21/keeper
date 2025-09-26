// cmd/api/main.go
// in cmd/api/main.go
package main

// @title           KEEPER API
// @version         1.0
// @description     API for the KEEPER Dealership Management System.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Gianni De Grossi
// @contact.url    https://github.com/GianniDG21
// @contact.email  giannidegrossi@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

import (
	"log"
	"os"
	"keeper/internal/api"
	"keeper/internal/storage"
	"github.com/joho/godotenv"
	"github.com/go-playground/validator/v10"
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