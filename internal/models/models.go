package models

type Dealership struct {
	ID_Dealership int `json:"id_dealership"`
	PostalCode    string `json:"postal_code"`
	City         string `json:"city"`
	Address      string `json:"address"`
	Phone       string `json:"phone"`
}

//Qui andranno le altre strutture dati (models) che rappresentano le entità del nostro dominio.