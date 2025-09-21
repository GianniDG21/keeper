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
	ID_Employee int    `json:"id_employee"`
	Role        Role   `json:"role"`
	TIN         string `json:"tin"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Phone       string `json:"phone"`
}

type Employment struct {
	ID_Employment int     `json:"id_employment"`
	ID_Employee   int     `json:"id_employee"`
	ID_Dealership int     `json:"id_dealership"`
	StartDate    time.Time `json:"start_date"`
	EndDate       *time.Time `json:"end_date,omitempty"` //Pointer to allow null values
}

type ClientType string //ClientType as Enum for Client struct
const (
	ClientTypeIndividual ClientType = "Private"
	ClientTypeCompany    ClientType = "Business"
)

type Client struct {
	ID_Client  int        `json:"id_client"`
	Type       ClientType `json:"type"`
	Phone      string     `json:"phone"`
	TIN_VAT    string     `json:"tin_vat"`
	Name       string     `json:"name"`
	Surname    *string    `json:"surname,omitempty"`    //Pointer to allow null values
	Company    *string    `json:"company,omitempty"`    //Pointer to allow null values
	Profession *string    `json:"profession,omitempty"` //Pointer to allow null values
}

type CondType string //CondType as Enum for CarPark struct
const (
	CondTypeNew  CondType = "new"
	CondTypeUsed CondType = "used"
)

type CarPark struct {
	VIN           string   `json:"vin"`
	ID_Dealership int      `json:"id_dealership"`
	Brand         string   `json:"brand"`
	Model         string   `json:"model"`
	Condition     CondType `json:"condition"`
	Year          int   `json:"year"`
	KM            int   `json:"km"`
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