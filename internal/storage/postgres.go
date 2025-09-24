// internal/storage/postgres.go
package storage

import (
	"database/sql"
	"fmt"
	"keeper/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresStore struct {
	db     *sql.DB
	gormDB *gorm.DB
}

func checkResult(result *gorm.DB) error { // This function checks the result of a GORM operation
	// Check for generic GORM errors
	if result.Error != nil {
		return result.Error
	}

	// Check if no rows were affected (ID not found)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	// If everything went well, there is no error
	return nil
}

func NewPostgresStore(connString string) (*PostgresStore, error) {
	// 1. Apriamo la connessione principale con GORM
	gormDB, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 2. Estraiamo la connessione *sql.DB sottostante da GORM
	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, err
	}

	// 3. Facciamo un Ping per verificare che la connessione standard sia viva
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	log.Println("Database connected successfully (GORM & standard SQL)")

	// 4. Restituiamo uno store con entrambi i campi popolati
	return &PostgresStore{db: sqlDB, gormDB: gormDB}, nil
}

// SQL Queries

// DEALERSHIP QUERIES
func (s *PostgresStore) CreateDealership(dealership *models.Dealership) (int, error) {
	query := `INSERT INTO dealership (postalcode, city, address, phone) 
			  VALUES ($1, $2, $3, $4) 
			  RETURNING id_dealership`

	var newID int
	// MANCAVA QUESTO .Scan(&newID)
	err := s.db.QueryRow(
		query,
		dealership.PostalCode,
		dealership.City,
		dealership.Address,
		dealership.Phone,
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

func (s *PostgresStore) GetDealerships() ([]*models.Dealership, error) {
	query := `SELECT id_dealership, postalcode, city, address, phone FROM dealership`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dealerships []*models.Dealership
	for rows.Next() {
		dealership := new(models.Dealership)
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
              SET postalcode = $1, city = $2, address = $3, phone = $4 
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
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("nessuna concessionaria trovata con id %d", id)
	}

	return nil
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

// EMPLOYEE QUERIES
func (s *PostgresStore) CreateEmployee(employee *models.Employee) (int, error) {
	result := s.gormDB.Create(employee)
	return employee.ID_Employee, result.Error
}

func (s *PostgresStore) GetEmployee() ([]*models.Employee, error) {
	var employee []*models.Employee
	result := s.gormDB.Find(&employee)
	return employee, result.Error
}

func (s *PostgresStore) UpdateEmployee(id int, employee *models.Employee) error {
	employee.ID_Employee = id
	result := s.gormDB.Save(employee)
	return checkResult(result)
}

func (s *PostgresStore) DeleteEmployee(id int) error {
	result := s.gormDB.Delete(&models.Employee{}, id)
	return checkResult(result)
}

// EMPLOYMENT QUERIES
func (s *PostgresStore) CreateEmployment(employment *models.Employment) (int, error) {
	result := s.gormDB.Create(employment)
	return employment.ID_Employment, result.Error
}

func (s *PostgresStore) GetEmployments() ([]*models.Employment, error) {
	var employments []*models.Employment
	result := s.gormDB.Find(&employments)
	return employments, result.Error
}

func (s *PostgresStore) UpdateEmployment(id int, employment *models.Employment) error {
	employment.ID_Employment = id
	result := s.gormDB.Save(employment)
	return checkResult(result)
}

func (s *PostgresStore) DeleteEmployment(id int) error {
	result := s.gormDB.Delete(&models.Employment{}, id)
	return checkResult(result)
}

// CLIENT QUERIES
func (s *PostgresStore) CreateClient(client *models.Client) (int, error) {
	result := s.gormDB.Create(client)
	return client.ID_Client, result.Error
}

func (s *PostgresStore) GetClients() ([]*models.Client, error) {
	var clients []*models.Client
	result := s.gormDB.Find(&clients)
	return clients, result.Error
}

func (s *PostgresStore) UpdateClient(id int, client *models.Client) error {
	client.ID_Client = id
	result := s.gormDB.Save(client)
	return checkResult(result)
}

func (s *PostgresStore) DeleteClient(id int) error {
	result := s.gormDB.Delete(&models.Client{}, id)
	return checkResult(result)
}

// CARPARK QUERIES
func (s *PostgresStore) CreateCarPark(carPark *models.CarPark) (int, error) {
	result := s.gormDB.Create(carPark)
	return carPark.ID_Car, result.Error
}

func (s *PostgresStore) GetCarParks() ([]*models.CarPark, error) {
	var carParks []*models.CarPark
	result := s.gormDB.Find(&carParks)
	return carParks, result.Error
}

func (s *PostgresStore) UpdateCarPark(id int, carPark *models.CarPark) error {
	carPark.ID_Car = id
	result := s.gormDB.Save(carPark)
	return checkResult(result)
}

func (s *PostgresStore) DeleteCarPark(id int) error {
	result := s.gormDB.Delete(&models.CarPark{}, id)
	return checkResult(result)
}

// ORDER QUERIES
func (s *PostgresStore) CreateOrder(order *models.Order) (int, error) {
	result := s.gormDB.Create(order)
	return order.ID_Order, result.Error
}

func (s *PostgresStore) GetOrders() ([]*models.Order, error) {
	var orders []*models.Order
	result := s.gormDB.Find(&orders)
	return orders, result.Error
}

func (s *PostgresStore) UpdateOrder(id int, order *models.Order) error {
	order.ID_Order = id
	result := s.gormDB.Save(order)
	return checkResult(result)
}

func (s *PostgresStore) DeleteOrder(id int) error {
	result := s.gormDB.Delete(&models.Order{}, id)
	return checkResult(result)
}

// APPOINTMENT QUERIES
func (s *PostgresStore) CreateAppointment(appointment *models.Appointment) (int, error) {
	result := s.gormDB.Create(appointment)
	return appointment.ID_Appointment, result.Error
}

func (s *PostgresStore) GetAppointments() ([]*models.Appointment, error) {
	var appointments []*models.Appointment
	result := s.gormDB.Find(&appointments)
	return appointments, result.Error
}

func (s *PostgresStore) UpdateAppointment(id int, appointment *models.Appointment) error {
	appointment.ID_Appointment = id
	result := s.gormDB.Save(appointment)
	return checkResult(result)
}

func (s *PostgresStore) DeleteAppointment(id int) error {
	result := s.gormDB.Delete(&models.Appointment{}, id)
	return checkResult(result)
}
