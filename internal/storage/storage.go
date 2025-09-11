// internal/storage/storage.go
package storage

import "keeper/internal/models"

//PROVA
//PROVA
type Store interface {
	GetDealerships() ([]*models.Dealership, error)
	//Qui aggiungero' create vehicle ecc
}
