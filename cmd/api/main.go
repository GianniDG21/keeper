// cmd/api/main.go
package main

import (
	"log"
	"keeper/internal/storage" 
)

func main() {
	connString := "user=keeper password=keeper dbname=keeper sslmode=disable"
	
	_, err := storage.NewStore(connString)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Setup completed. Server will start here...")
}