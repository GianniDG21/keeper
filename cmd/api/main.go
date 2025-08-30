// cmd/api/main.go
package main

import (
	"gianni/keeper/internal/storage" // Assicurati di usare il path corretto del tuo modulo
	"log"
)

func main() {
	// Stringa di connessione per il nostro database Docker
	connString := "user=keeper password=keeper dbname=keeper sslmode=disable"

	// Inizializza lo store (e quindi la connessione al DB)
	_, err := storage.NewStore(connString)
	if err != nil {
		log.Fatal(err)
	}

	// Per ora terminiamo qui, in seguito avvieremo il server.
	log.Println("Setup completed. Server will start here...")
}
