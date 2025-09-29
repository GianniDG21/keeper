// internal/storage/postgres.go
package storage

import (
	"database/sql"
	"fmt"
	"keeper/internal/models"
	"log"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresStore struct {
	Db     *sql.DB
	GormDB *gorm.DB
}

type dependencyCheck struct {
	model     interface{}
	fieldName string
}

func (s *PostgresStore) checkDependencies(parentID int, checks map[string]dependencyCheck) error {
	var errorMessages []string

	for parentName, check := range checks {
		var count int64
		query := fmt.Sprintf("%s = ?", check.fieldName)
		
		s.GormDB.Model(check.model).Where(query, parentID).Count(&count)
		
		if count > 0 {
			errorMessages = append(errorMessages, fmt.Sprintf("referenced by %d %s records", count, parentName))
		}
	}

	if len(errorMessages) > 0 {
		return fmt.Errorf("cannot delete: %s", strings.Join(errorMessages, ", "))
	}

	return nil
}

func checkResult(result *gorm.DB) error {
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func NewPostgresStore(connString string) (*PostgresStore, error) {
	gormDB, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	log.Println("Database connected successfully (GORM & standard SQL)")

	return &PostgresStore{Db: sqlDB, GormDB: gormDB}, nil
}

func (s *PostgresStore) CreateDealership(dealership *models.Dealership) (int, error) {
	query := `INSERT INTO dealership (postalcode, city, address, phone) 
			  VALUES ($1, $2, $3, $4) 
			  RETURNING id_dealership`

	var newID int
	err := s.Db.QueryRow(
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
	rows, err := s.Db.Query(query)
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

	result, err := s.Db.Exec(
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
		return fmt.Errorf("no dealership found with id %d", id)
	}

	return nil
}

func (s *PostgresStore) DeleteDealership(id int) error {
	checks := map[string]dependencyCheck{
		"cars":         {&models.CarPark{}, "id_dealership"},
		"employments":  {&models.Employment{}, "id_dealership"},
		"orders":       {&models.Order{}, "id_dealership"},
		"appointments": {&models.Appointment{}, "id_dealership"},
	}
	
	if err := s.checkDependencies(id, checks); err != nil {
		return err
	}

	query := `DELETE FROM dealership WHERE id_dealership = $1`
	result, err := s.Db.Exec(query, id)
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

func (s *PostgresStore) CreateEmployee(employee *models.Employee) (int, error) {
	result := s.GormDB.Create(employee)
	if result.Error != nil {
		return 0, result.Error  
	}
	return employee.ID_Employee, nil
}

func (s *PostgresStore) GetEmployees() ([]*models.Employee, error) {
	var employees []*models.Employee 
	result := s.GormDB.Find(&employees)
	return employees, result.Error
}

func (s *PostgresStore) UpdateEmployee(id int, employee *models.Employee) error {
	employee.ID_Employee = id
	result := s.GormDB.Save(employee)
	return checkResult(result)
}

func (s *PostgresStore) DeleteEmployee(id int) error {
	checks := map[string]dependencyCheck{
		"orders":       {&models.Order{}, "id_employee"},
		"appointments": {&models.Appointment{}, "id_employee"},
		"employments":  {&models.Employment{}, "id_employee"},
	}
	
	if err := s.checkDependencies(id, checks); err != nil {
		return err
	}

	result := s.GormDB.Delete(&models.Employee{}, id)
	return checkResult(result)
}

func (s *PostgresStore) CreateEmployment(employment *models.Employment) (int, error) {
	result := s.GormDB.Create(employment)
	if result.Error != nil {
		return 0, result.Error
	}
	return employment.ID_Employment, nil
}

func (s *PostgresStore) GetEmployments() ([]*models.Employment, error) {
	var employments []*models.Employment
	result := s.GormDB.Find(&employments)
	return employments, result.Error
}

func (s *PostgresStore) UpdateEmployment(id int, employment *models.Employment) error {
	employment.ID_Employment = id
	result := s.GormDB.Save(employment)
	return checkResult(result)
}

func (s *PostgresStore) DeleteEmployment(id int) error {
	result := s.GormDB.Delete(&models.Employment{}, id)
	return checkResult(result)
}

func (s *PostgresStore) CreateClient(client *models.Client) (int, error) {
	result := s.GormDB.Create(client)
	if result.Error != nil {
		return 0, result.Error
	}
	return client.ID_Client, nil
}

func (s *PostgresStore) GetClients() ([]*models.Client, error) {
	var clients []*models.Client
	result := s.GormDB.Find(&clients)
	return clients, result.Error
}

func (s *PostgresStore) UpdateClient(id int, client *models.Client) error {
	client.ID_Client = id
	result := s.GormDB.Save(client)
	return checkResult(result)
}

func (s *PostgresStore) DeleteClient(id int) error {
	checks := map[string]dependencyCheck{
		"orders":       {&models.Order{}, "id_client"},
		"appointments": {&models.Appointment{}, "id_client"},
	}
	
	if err := s.checkDependencies(id, checks); err != nil {
		return err
	}

	result := s.GormDB.Delete(&models.Client{}, id)
	return checkResult(result)
}

func (s *PostgresStore) CreateCar(car *models.CarPark) (int, error) {
	result := s.GormDB.Create(car)
	if result.Error != nil {
		return 0, result.Error
	}
	return car.ID_Car, nil
}

func (s *PostgresStore) GetCars() ([]*models.CarPark, error) {
	var carParks []*models.CarPark
	result := s.GormDB.Find(&carParks)
	return carParks, result.Error
}

func (s *PostgresStore) PatchCar(id int, updates map[string]interface{}) error {
	result := s.GormDB.Model(&models.CarPark{}).Where("id_car = ?", id).Updates(updates)
	return checkResult(result)
}

func (s *PostgresStore) DeleteCar(id int) error {
	var car models.CarPark
	if err := s.GormDB.First(&car, id).Error; err != nil {
		return err
	}
	
	if car.VIN != nil {
		var count int64
		s.GormDB.Model(&models.Order{}).Where("vin = ?", *car.VIN).Count(&count)
		
		if count > 0 {
			return fmt.Errorf("cannot delete car: referenced by %d order records", count)
		}
	}
	
	result := s.GormDB.Delete(&models.CarPark{}, id)
	return checkResult(result)
}

func (s *PostgresStore) CreateOrder(order *models.Order) (int, error) {
	result := s.GormDB.Create(order)
	if result.Error != nil {
		return 0, result.Error
	}
	return order.ID_Order, nil
}

func (s *PostgresStore) GetOrders() ([]*models.Order, error) {
	var orders []*models.Order
	result := s.GormDB.Find(&orders)
	return orders, result.Error
}

func (s *PostgresStore) UpdateOrder(id int, order *models.Order) error {
	order.ID_Order = id
	result := s.GormDB.Save(order)
	return checkResult(result)
}

func (s *PostgresStore) DeleteOrder(id int) error {
	result := s.GormDB.Delete(&models.Order{}, id)
	return checkResult(result)
}

func (s *PostgresStore) CreateAppointment(appointment *models.Appointment) (int, error) {
	result := s.GormDB.Create(appointment)
	if result.Error != nil {
		return 0, result.Error
	}
	return appointment.ID_Appointment, nil
}

func (s *PostgresStore) GetAppointments() ([]*models.Appointment, error) {
	var appointments []*models.Appointment
	result := s.GormDB.Find(&appointments)
	return appointments, result.Error
}

func (s *PostgresStore) UpdateAppointment(id int, appointment *models.Appointment) error {
	appointment.ID_Appointment = id
	result := s.GormDB.Save(appointment)
	return checkResult(result)
}

func (s *PostgresStore) DeleteAppointment(id int) error {
	result := s.GormDB.Delete(&models.Appointment{}, id)
	return checkResult(result)
}
