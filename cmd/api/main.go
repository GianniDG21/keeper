// cmd/api/main.go
package main

import (
	"log"
	"os"
	"keeper/internal/api"
	"keeper/internal/storage"
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
	
	// Usiamo la nuova funzione NewPostgresStore
	store, err := storage.NewPostgresStore(connString)
	if err != nil {
		log.Fatal("failed to connect to the database: ", err)
	}

	server := api.NewAPIServer(":8080", store)
	server.Run()
}