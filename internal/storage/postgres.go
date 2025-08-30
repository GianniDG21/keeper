// internal/storage/postgres.go
package storage

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib" // Il driver viene importato per i suoi "side effects"
)

// Store gestirà tutte le interazioni con il database
type Store struct {
	db *sql.DB
}

// NewStore crea una nuova istanza dello Store e stabilisce la connessione al DB
func NewStore(connString string) (*Store, error) {
	db, err := sql.Open("pgx", connString)
	if err != nil {
		return nil, err
	}

	// Verifica che la connessione sia viva
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Database connected successfully")

	return &Store{db: db}, nil
}
