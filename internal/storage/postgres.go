// internal/storage/postgres.go
package storage

import (
	"database/sql"
	"keeper/internal/models"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresStore struct {
	db *sql.DB
}

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

// SQL Queries

// DEALERSHIP QUERIES
func (s *PostgresStore) CreateDealership() ([]models.Dealership, error) {
	query := `INSERT INTO Dealership (postal_code, city, address, phone) VALUES ($1, $2, $3, $4) RETURNING id_dealership`

	

func (s *PostgresStore) GetAllDealerships() ([]models.Dealership, error) {
	query := `SELECT id_dealership, postal_code, city, address, phone FROM Dealership`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dealerships []models.Dealership
	for rows.Next() {
		var dealership models.Dealership
		err := rows.Scan(
			&dealership.ID_Dealership,
			&dealership.PostalCode,
			&dealership.City,
			&dealership.Address,
			&dealership.Phone,
		)
		if err != nil {
			return nil, err
		}
		dealerships = append(dealerships, dealership)
	}

	return dealerships, nil
}
