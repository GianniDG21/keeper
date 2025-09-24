// internal/api/server.go
package api

import (
	"keeper/internal/storage" // Assicurati che il path del modulo sia corretto
	"log"
	"net/http"
)

type APIServer struct {
	listenAddr string
	store      storage.Store
}

func NewAPIServer(listenAddr string, store storage.Store) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

// Run avvia il server HTTP e registra tutte le rotte (gli endpoint).
func (s *APIServer) Run() {
	router := http.NewServeMux()

	router.HandleFunc("GET /healthcheck", s.handleHealthCheck)


	// Dealerships endpoints //
	router.HandleFunc("POST /dealerships", s.handleCreateDealership)
	router.HandleFunc("GET /dealerships", s.handleGetDealerships)
	router.HandleFunc("PUT /dealerships/{id}", s.handleUpdateDealership)
	router.HandleFunc("DELETE /dealerships/{id}", s.handleDeleteDealership)

	// Employee endpoints //
	router.HandleFunc("POST /employees", s.handleCreateEmployee)
	router.HandleFunc("GET /employees", s.handleGetEmployee)
	router.HandleFunc("PUT /employees/{id}", s.handleUpdateEmployee)
	router.HandleFunc("DELETE /employees/{id}", s.handleDeleteEmployee)

	// Employment endpoints //
	router.HandleFunc("POST /employments", s.handleCreateEmployment)
	router.HandleFunc("GET /employments", s.handleGetEmployments)
	router.HandleFunc("PUT /employments/{id}", s.handleUpdateEmployment)
	router.HandleFunc("DELETE /employments/{id}", s.handleDeleteEmployment)

	
	log.Println("JSON API server running on port", s.listenAddr)

	err := http.ListenAndServe(s.listenAddr, router)
	if err != nil {
		log.Fatal(err)
	}
}
