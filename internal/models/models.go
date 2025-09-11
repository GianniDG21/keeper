package models

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
	StartDate     string  `json:"start_date"`
	EndDate       *string `json:"end_date,omitempty"` //Pointer to allow null values
}

type ClientType string //ClientType as Enum for Client struct
const (
	ClientTypeIndividual ClientType = "Private"
	ClientTypeCompany    ClientType = "Business"
)

type Client struct {
	ID_Client  int        `json:"id_client"`
	Type       ClientType `json:"type"`
	TIN_VAT    string     `json:"tin_vat"`
	Name       string     `json:"name"`
	Surname    *string    `json:"surname,omitempty"` //Pointer to allow null values
	Phone      string     `json:"phone"`
	Email      string     `json:"email"`
	Profession *string    `json:"profession,omitempty"` //Pointer to allow null values
}

type CarPark struct {
	VIN string `json:"vin"`
	//INFORMARSI SU COME FUNZIONANO LE CHIAVI SECONDARIE
}

type OrderStatus string //OrderStatus as Enum for Order struct
const (
	OrderStatusClient     OrderStatus = "client_pending"
	OrderStatusCompleted  OrderStatus = "completed"
	OrderStatusCancelled  OrderStatus = "cancelled"
	OrderStatusInProgress OrderStatus = "in_progress"
)

//type Order
//Qui andranno le altre strutture dati (models) che rappresentano le entità del nostro dominio.
