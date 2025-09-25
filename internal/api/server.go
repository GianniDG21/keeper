// internal/api/server.go
package api

import (
	"keeper/internal/storage"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"          // <-- NUOVO IMPORT
	"github.com/go-chi/chi/v5/middleware" // <-- NUOVO IMPORT
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

func (s *APIServer) Run() {
	// 1. Creiamo un nuovo router usando chi
	router := chi.NewRouter()

	// 2. Aggiungiamo alcuni "middleware" standard e utili
	router.Use(middleware.Logger)    // Logga ogni richiesta in arrivo
	router.Use(middleware.Recoverer) // Gestisce eventuali "panic" e impedisce al server di crashare

	// 3. Registriamo le rotte raggruppandole logicamente

	router.Get("/healthcheck", s.handleHealthCheck) // GET /healthcheck

	router.Route("/dealerships", func(r chi.Router) {
		r.Post("/", s.handleCreateDealership)   // POST /dealerships
		r.Get("/", s.handleGetDealerships)      // GET /dealerships
		r.Put("/{id}", s.handleUpdateDealership)  // PUT /dealerships/{id}
		r.Delete("/{id}", s.handleDeleteDealership) // DELETE /dealerships/{id}
	})

	router.Route("/employees", func(r chi.Router) {
		r.Post("/", s.handleCreateEmployee)   // POST /employees
		r.Get("/", s.handleGetEmployee)      // GET /employees
		r.Put("/{id}", s.handleUpdateEmployee)  // PUT /employees/{id}
		r.Delete("/{id}", s.handleDeleteEmployee) // DELETE /employees/{id}
	})

	router.Route("/employments", func(r chi.Router) {
		r.Post("/", s.handleCreateEmployment)   // POST /employments
		r.Get("/", s.handleGetEmployments)      // GET /employments
		r.Put("/{id}", s.handleUpdateEmployment)  // PUT /employments/{id}
		r.Delete("/{id}", s.handleDeleteEmployment) // DELETE /employments/{id}
	})

	router.Route("/clients", func(r chi.Router) {
		r.Post("/", s.handleCreateClient)   // POST /clients
		r.Get("/", s.handleGetClients)      // GET /clients
		r.Put("/{id}", s.handleUpdateClient)  // PUT /clients/{id}
		r.Delete("/{id}", s.handleDeleteClient) // DELETE /clients/{id}
	})

	log.Println("JSON API server running on port", s.listenAddr)

	// La chiamata per avviare il server rimane la stessa
	err := http.ListenAndServe(s.listenAddr, router)
	if err != nil {
		log.Fatal(err)
	}
}