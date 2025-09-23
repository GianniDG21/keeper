package models

import (
	"time"
)


type Dealership struct {
	ID_Dealership int    `json:"id_dealership"`
	PostalCode    string `json:"postal_code"`
	City          string `json:"city"`
	Address       string `json:"address"`
	Phone         string `json:"phone"`
}
type Role string //Role Type as Enum for Employee struct
const (
	RoleAssistant Role = "assistant"
	RoleSeller    Role = "seller"
	RoleManager   Role = "manager"
	RoleAdmin     Role = "admin"
	RoleMechanic  Role = "mechanic"
)

type Employee struct {
	ID_Employee int    `json:"id_employee" gorm:"primaryKey;autoIncrement"`
	Role        Role   `json:"role" gorm:"column:role"`
	TIN         string `json:"tin" gorm:"column:tin;unique"`
	Name        string `json:"name" gorm:"column:name"`
	Surname     string `json:"surname" gorm:"column:surname"`
	Phone       string `json:"phone" gorm:"column:phone"`
}

type Employment struct {
	ID_Employment int        `json:"id_employment" gorm:"primaryKey;autoIncrement"`
	ID_Employee   int        `json:"id_employee" gorm:"column:id_employee;not null"`
	ID_Dealership int        `json:"id_dealership" gorm:"column:id_dealership;not null"`
	StartDate     time.Time  `json:"start_date" gorm:"column:start_date;not null"`
	EndDate       *time.Time `json:"end_date,omitempty" gorm:"column:end_date"`
}

type ClientType string //ClientType as Enum for Client struct
const (
	ClientTypeIndividual ClientType = "Private"
	ClientTypeCompany    ClientType = "Business"
)

type Client struct {
	ID_Client  int        `json:"id_client" gorm:"primaryKey;autoIncrement"`
	Type       ClientType `json:"type" gorm:"column:type"`
	Phone      string     `json:"phone" gorm:"column:phone"`
	TIN_VAT    string     `json:"tin_vat" gorm:"column:tin_vat;unique"`
	Name       string     `json:"name" gorm:"column:name"`
	Surname    *string    `json:"surname,omitempty" gorm:"column:surname"`       //Pointer to allow null values
	Company    *string    `json:"company,omitempty" gorm:"column:company"`       //Pointer to allow null values
	Profession *string    `json:"profession,omitempty" gorm:"column:profession"` //Pointer to allow null values
}

type CondType string //CondType as Enum for CarPark struct
const (
	CondTypeNew  CondType = "new"
	CondTypeUsed CondType = "used"
)

type CarPark struct {
	ID_Car	      int      `json:"id_car" gorm:"primaryKey;autoIncrement"`
	VIN           string   `json:"vin" gorm:"primaryKey;column:vin"`
	ID_Dealership int      `json:"id_dealership" gorm:"column:id_dealership;not null"`
	Brand         string   `json:"brand" gorm:"column:brand"`
	Model         string   `json:"model" gorm:"column:model"`
	Condition     CondType `json:"condition" gorm:"column:condition"`
	Year          int      `json:"year" gorm:"column:year"`
	KM            int      `json:"km" gorm:"column:km"`
}

type OrderStatus string //OrderStatus as Enum for Order struct
const (
	OrderStatusClient     OrderStatus = "client_pending"
	OrderStatusCompleted  OrderStatus = "completed"
	OrderStatusCancelled  OrderStatus = "cancelled"
	OrderStatusInProgress OrderStatus = "in_progress"
)

type Order struct {
	ID_Order      int         `json:"id_order"`
	Status        OrderStatus `json:"status"`
	ID_Client     int         `json:"id_client"`
	ID_Employee   int         `json:"id_employee"`
	VIN           string      `json:"vin"`
	ID_Dealership int         `json:"id_dealership"`
	LastUpdate    time.Time        `json:"last_update"`

}

type Appointment struct {
	ID_Appointment int       `json:"id_appointment"`
	ID_Client      int       `json:"id_client"`
	ID_Employee    int       `json:"id_employee"`
	ID_Dealership  int       `json:"id_dealership"`
	Date           time.Time `json:"date"`
	Reason		 string    `json:"reason"`
	Note		  *string   `json:"note,omitempty"` //Pointer to allow null values
}