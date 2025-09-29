package models

import (
	"encoding/json"
	"time"
)

type Dealership struct {
	ID_Dealership int    `json:"id_dealership" gorm:"primaryKey;autoIncrement"`
	PostalCode    string `json:"postalcode" gorm:"column:postalcode;not null" validate:"required,max=5"`
	City          string `json:"city" gorm:"column:city;not null" validate:"required,max=30"`
	Address       string `json:"address" gorm:"column:address;not null" validate:"required,max=100"`
	Phone         string `json:"phone" gorm:"column:phone;not null" validate:"required,max=20"`
}

type Role string
const (
	RoleAssistant   Role = "assistant"
	RoleSalesperson Role = "salesperson"
	RoleManager     Role = "manager"
	RoleAdmin       Role = "admin"
	RoleMechanic    Role = "mechanic"
)

type Employee struct {
	ID_Employee int    `json:"id_employee" gorm:"primaryKey;autoIncrement"`
	Role        Role   `json:"role" gorm:"column:role;not null;default:assistant" validate:"required,oneof=assistant salesperson manager admin mechanic"`
	TIN         string `json:"tin" gorm:"column:tin;unique;not null" validate:"required,max=16"`
	Name        string `json:"name" gorm:"column:name;not null" validate:"required,max=50"`
	Surname     string `json:"surname" gorm:"column:surname;not null" validate:"required,max=50"`
	Phone       string `json:"phone" gorm:"column:phone;not null" validate:"required,max=20"`
}

type Employment struct {
	ID_Employment int        `json:"id_employment" gorm:"primaryKey;autoIncrement"`
	ID_Employee   int        `json:"id_employee" gorm:"column:id_employee;not null" validate:"required"`
	ID_Dealership int        `json:"id_dealership" gorm:"column:id_dealership;not null" validate:"required"`
	StartDate     time.Time  `json:"startdate" gorm:"column:startdate;not null" validate:"required"`
	EndDate       *time.Time `json:"enddate,omitempty" gorm:"column:enddate"`
}

func (e *Employment) UnmarshalJSON(data []byte) error {
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

type ClientType string
const (
	ClientTypePrivate ClientType = "private"
	ClientTypeCompany ClientType = "company"
)

type Client struct {
	ID_Client   int         `json:"id_client" gorm:"primaryKey;autoIncrement"`
	Type        ClientType  `json:"type" gorm:"column:type;not null" validate:"required,oneof=private company"`
	Phone       *string     `json:"phone,omitempty" gorm:"column:phone" validate:"omitempty,max=20"`
	Email       *string     `json:"email,omitempty" gorm:"column:email;unique" validate:"omitempty,email,max=50"`
	TIN_VAT     string      `json:"tin_vat" gorm:"column:tin_vat;unique;not null" validate:"required,max=16"`
	Name        string      `json:"name" gorm:"column:name;not null" validate:"required,max=50"`
	Surname     *string     `json:"surname,omitempty" gorm:"column:surname" validate:"omitempty,max=50"`
	CompanyName *string     `json:"companyname,omitempty" gorm:"column:companyname" validate:"omitempty,max=100"`
	Profession  *string     `json:"profession,omitempty" gorm:"column:profession" validate:"omitempty,max=50"`
}

type CondType string
const (
	CondTypeNew  CondType = "new"
	CondTypeUsed CondType = "used"
)

type CarPark struct {
	ID_Car        int      `json:"id_car" gorm:"primaryKey;autoIncrement"`
	VIN           *string  `json:"vin,omitempty" gorm:"column:vin;unique" validate:"omitempty,alphanum,len=17"`
	ID_Dealership int      `json:"id_dealership" gorm:"column:id_dealership;not null" validate:"required"`
	Brand         string   `json:"brand" gorm:"column:brand;not null" validate:"required,max=30"`
	Model         string   `json:"model" gorm:"column:model;not null" validate:"required,max=30"`
	Condition     CondType `json:"condition" gorm:"column:condition;not null;default:new" validate:"required,oneof=new used"`
	Year          int      `json:"year" gorm:"column:year;not null" validate:"required,min=1901"`
	KM            string   `json:"km" gorm:"column:km;not null;default:'0'" validate:"required,max=7"`
	Plate         string   `json:"plate" gorm:"column:plate;unique;not null" validate:"required,max=10"`
}

type OrderStatus string
const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusCompleted  OrderStatus = "completed"
	OrderStatusCancelled  OrderStatus = "cancelled"
	OrderStatusInProgress OrderStatus = "in_progress"
)

type Order struct {
	ID_Order      int         `json:"id_order" gorm:"primaryKey;autoIncrement"`
	Status        OrderStatus `json:"status" gorm:"column:status;not null;default:pending" validate:"required,oneof=pending completed cancelled in_progress"`
	ID_Client     int         `json:"id_client" gorm:"column:id_client;not null" validate:"required"`
	ID_Employee   int         `json:"id_employee" gorm:"column:id_employee;not null" validate:"required"`
	VIN           string      `json:"vin" gorm:"column:vin;not null" validate:"required,alphanum,len=17"`
	ID_Dealership int         `json:"id_dealership" gorm:"column:id_dealership;not null" validate:"required"`
	LastUpdate    time.Time   `json:"last_update" gorm:"column:last_update;not null;default:CURRENT_TIMESTAMP"`
}

type Appointment struct {
	ID_Appointment int       `json:"id_appointment" gorm:"primaryKey;autoIncrement"`
	ID_Client      int       `json:"id_client" gorm:"column:id_client;not null" validate:"required"`
	ID_Employee    int       `json:"id_employee" gorm:"column:id_employee;not null" validate:"required"`
	ID_Dealership  int       `json:"id_dealership" gorm:"column:id_dealership;not null" validate:"required"`
	Date           time.Time `json:"date" gorm:"column:date;not null" validate:"required"`
	Reason         string    `json:"reason" gorm:"column:reason;not null" validate:"required,max=100"`
	Notes          *string   `json:"notes,omitempty" gorm:"column:notes"`
}

func (Employee) TableName() string {
	return "employee"
}
func (Employment) TableName() string {
	return "employment"
}
func (Client) TableName() string {
	return "client"
}
func (CarPark) TableName() string {
	return "car_park"
}
func (Order) TableName() string {
	return "order"
}
func (Appointment) TableName() string {
	return "appointment"
}
func (Dealership) TableName() string {
	return "dealership"
}
