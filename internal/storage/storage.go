// internal/storage/storage.go
package storage

import "keeper/internal/models"


type Store interface {
	//-----Dealership Methods-----
	CreateDealership(dealership *models.Dealership) (int, error)
	GetDealerships() ([]*models.Dealership, error)
	UpdateDealership(id int, dealership *models.Dealership) error
	DeleteDealership(id int) error

	//-----Employee Methods-----
	CreateEmployee(employee *models.Employee) (int, error)
	GetEmployees() ([]*models.Employee, error)
	UpdateEmployee(id int, employee *models.Employee) error
	DeleteEmployee(id int) error

	//-----Employment Methods-----
	CreateEmployment(employment *models.Employment) (int, error)
	GetEmployments() ([]*models.Employment, error)
	UpdateEmployment(id int, employment *models.Employment) error
	DeleteEmployment(id int) error

	//-----Client Methods-----
	CreateClient(client *models.Client) (int, error)
	GetClients() ([]*models.Client, error)
	UpdateClient(id int, client *models.Client) error
	DeleteClient(id int) error

	//-----CarPark Methods-----
	CreateCarPark(carPark *models.CarPark) (int, error)
	GetCarParks() ([]*models.CarPark, error)
	UpdateCarPark(id int, carPark *models.CarPark) error
	DeleteCarPark(vin string) error

	//-----Order Methods-----
	CreateOrder(order *models.Order) (int, error)
	GetOrders() ([]*models.Order, error)
	UpdateOrder(id int, order *models.Order) error
	DeleteOrder(id int) error

	//-----Appointment Methods-----
	CreateAppointment(appointment *models.Appointment) (int, error)
	GetAppointments() ([]*models.Appointment, error)
	UpdateAppointment(id int, appointment *models.Appointment) error
	DeleteAppointment(id int) error	
}
