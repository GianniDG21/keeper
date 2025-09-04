// cmd/api/main.go
package main

import (
	"keeper/internal/api"
	"keeper/internal/storage"
	"log"
)

func main() {
	// Inseriamo la stringa di connessione direttamente nel codice per lo sviluppo locale.
	// Questa è la stringa per il tuo database Docker.
	connString := "postgres://keeper:keeper@localhost:5432/keeper?sslmode=disable"

	store, err := storage.NewPostgresStore(connString)
	if err != nil {
		log.Fatal("failed to connect to the database: ", err)
	}

	server := api.NewAPIServer(":8080", store)
	server.Run()
}
