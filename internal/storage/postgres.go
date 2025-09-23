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
func (s *PostgresStore) CreateDealership(dealership *models.Dealership) (int, error) {
	query := `INSERT INTO dealership (postal_code, city, address, phone) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id_dealership`

	var newID int

	err := s.db.QueryRow(
		query,
		dealership.PostalCode,
		dealership.City,
		dealership.Address,
		dealership.Phone,
	)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

func (s *PostgresStore) GetAllDealerships() ([]models.Dealership, error) {
	query := `SELECT id_dealership, postal_code, city, address, phone FROM dealership`

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

func (s *PostgresStore) UpdateDealership(id int, dealership *models.Dealership) error {
    query := `UPDATE dealership 
              SET postal_code = $1, city = $2, address = $3, phone = $4 
              WHERE id_dealership = $5`
    
    result, err := s.db.Exec(
        query,
        dealership.PostalCode,
        dealership.City,
        dealership.Address,
        dealership.Phone,
        id,
    )
    if err != nil {
        return err // Corretto: restituisce solo l'errore
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err // Corretto: restituisce solo l'errore
    }

    if rowsAffected == 0 {
        return fmt.Errorf("nessuna concessionaria trovata con id %d", id)
    }

    return nil // Corretto: in caso di successo, restituisce nil (nessun errore)
}

func (s *PostgresStore) DeleteDealership(id int) error {
	query := `DELETE FROM dealership WHERE id_dealership = $1`
	result, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// From now on, we will be using GORM for ORM operations
// this will simplify the code, improve security, and enhance maintainability
// I've decided to do Dealership's methods as an example to have a solid base

