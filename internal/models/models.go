package models

import (
	"time"
	"encoding/json"
)


type Dealership struct {
	ID_Dealership int    `json:"id_dealership"`
	PostalCode    string `json:"postal_code" validate:"len=5,numeric"`
	City          string `json:"city" validate:"required"`
	Address       string `json:"address" validate:"required"`
	Phone         string `json:"phone" validate:"required"`
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
	Role        Role   `json:"role" gorm:"column:role" validate:"required,oneof=assistant seller manager admin mechanic"`
	TIN         string `json:"tin" gorm:"column:tin;unique" validate:"required"`
	Name        string `json:"name" gorm:"column:name" validate:"required"`
	Surname     string `json:"surname" gorm:"column:surname" validate:"required"`
	Phone       string `json:"phone" gorm:"column:phone" validate:"required"`
}

type Employment struct {
	ID_Employment int        `json:"id_employment" gorm:"primaryKey;autoIncrement"`
	ID_Employee   int        `json:"id_employee" gorm:"column:id_employee;not null" validate:"required"`
	ID_Dealership int        `json:"id_dealership" gorm:"column:id_dealership;not null" validate:"required"`
	StartDate     time.Time  `json:"startdate" gorm:"column:startdate;not null" validate:"required"`
	EndDate       *time.Time `json:"enddate,omitempty" gorm:"column:enddate"`
}
func (e *Employment) UnmarshalJSON(data []byte) error {	// Custom UnmarshalJSON to handle date parsing
	type Alias Employment
	aux := &struct {
		StartDate string  `json:"startdate"`
		EndDate   *string `json:"enddate"`
		*Alias
	}{
		Alias: (*Alias)(e),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	startDate, err := time.Parse("2006-01-02", aux.StartDate)
	if err != nil {
		return err
	}
	e.StartDate = startDate 

	if aux.EndDate != nil {
		endDate, err := time.Parse("2006-01-02", *aux.EndDate)
		if err != nil {
			return err
		}
		e.EndDate = &endDate
	}

	return nil
}

type ClientType string //ClientType as Enum for Client struct
const (
	ClientTypeIndividual ClientType = "private"
	ClientTypeCompany    ClientType = "business"
)

type Client struct {
	ID_Client  int        `json:"id_client" gorm:"primaryKey;autoIncrement"`
	Type       ClientType `json:"type" gorm:"column:type" validate:"required,oneof=Private Business"`
	Phone      string     `json:"phone" gorm:"column:phone" validate:"required"`
	TIN_VAT    string     `json:"tin_vat" gorm:"column:tin_vat;unique" validate:"required"`
	Name       string     `json:"name" gorm:"column:name" validate:"required"`
	Surname    *string    `json:"surname,omitempty" gorm:"column:surname"`       //Pointer to allow null values
	Company    *string    `json:"companyname,omitempty" gorm:"column:companyname"`       //Pointer to allow null values
	Profession *string    `json:"profession,omitempty" gorm:"column:profession"` //Pointer to allow null values
}

type CondType string //CondType as Enum for CarPark struct
const (
	CondTypeNew  CondType = "new"
	CondTypeUsed CondType = "used"
)

type CarPark struct {
	ID_Car	      int      `json:"id_car" gorm:"primaryKey;autoIncrement"`
	VIN           string   `json:"vin" gorm:"primaryKey;column:vin" validate:"required,alphanum,len=17"`
	ID_Dealership int      `json:"id_dealership" gorm:"column:id_dealership;not null"`
	Brand         string   `json:"brand" gorm:"column:brand" validate:"required"`
	Model         string   `json:"model" gorm:"column:model" validate:"required"`
	Condition     CondType `json:"condition" gorm:"column:condition" validate:"required,oneof=new used"`
	Year          int      `json:"year" gorm:"column:year" validate:"required,min=1886,max=2023"`
	KM            int      `json:"km" gorm:"column:km"`
	Plate		 string   `json:"plate" gorm:"column:plate;unique"`
}

type OrderStatus string //OrderStatus as Enum for Order struct
const (
	OrderStatusClient     OrderStatus = "client_pending"
	OrderStatusCompleted  OrderStatus = "completed"
	OrderStatusCancelled  OrderStatus = "cancelled"
	OrderStatusInProgress OrderStatus = "in_progress"
)

type Order struct {
	ID_Order      int         `json:"id_order" gorm:"primaryKey;autoIncrement"`
	Status        OrderStatus `json:"status" gorm:"column:status" validate:"required,oneof=client_pending completed cancelled in_progress"`
	ID_Client     int         `json:"id_client" gorm:"column:id_client;not null" validate:"required"`
	ID_Employee   int         `json:"id_employee" gorm:"column:id_employee;not null" validate:"required"`
	VIN           string      `json:"vin" gorm:"column:vin;not null" validate:"required,alphanum,len=17"`
	ID_Dealership int         `json:"id_dealership" gorm:"column:id_dealership;not null" validate:"required"`
	LastUpdate    time.Time   `json:"last_update" gorm:"column:last_update;not null"`
}

type Appointment struct {
	ID_Appointment int       `json:"id_appointment" gorm:"primaryKey;autoIncrement"`
	ID_Client      int       `json:"id_client" gorm:"column:id_client;not null" validate:"required"`
	ID_Employee    int       `json:"id_employee" gorm:"column:id_employee;not null" validate:"required"`
	ID_Dealership  int       `json:"id_dealership" gorm:"column:id_dealership;not null" validate:"required"`
	Date           time.Time `json:"date" gorm:"column:date;not null" validate:"required"`
	Reason         string    `json:"reason" gorm:"column:reason" validate:"required"`
	Note           *string   `json:"note,omitempty" gorm:"column:note"`
}

// TableName overrides the default table name for some structs
func (Employee) TableName() string {
  return "employee"
}
func (Employment) TableName() string {
  return "employment"
}
func (Client) TableName() string {
	return "client"
}