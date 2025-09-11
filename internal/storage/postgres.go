// internal/storage/postgres.go
package storage

import (
	"database/sql"
	"keeper/internal/models"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// PostgresStore è la nostra implementazione concreta dell'interfaccia Store.
type PostgresStore struct {
	db *sql.DB
}

// NewPostgresStore crea una nuova istanza di PostgresStore.
func NewPostgresStore(connString string) (*PostgresStore, error) {
	db, err := sql.Open("pgx", connString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Database connected successfully")

	return &PostgresStore{db: db}, nil
}

// GetDealerships è il primo metodo che implementa parte dell'interfaccia Store.
func (s *PostgresStore) GetDealerships() ([]*models.Dealership, error) {
	rows, err := s.db.Query(`SELECT id_dealership, postal_code, city, address, phone FROM dealership`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dealerships []*models.Dealership
	for rows.Next() {
		dealership := new(models.Dealership)
		if err := rows.Scan(
			&dealership.ID_Dealership,
			&dealership.PostalCode,
			&dealership.City,
			&dealership.Address,
			&dealership.Phone,
		); err != nil {
			return nil, err
		}
		dealerships = append(dealerships, dealership)
	}

	return dealerships, nil
}