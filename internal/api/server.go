// internal/api/server.go
package api

import (
	"keeper/internal/storage"
	"log"
	"net/http"
	"github.com/go-playground/validator/v10"
	"github.com/go-chi/chi/v5"          
	"github.com/go-chi/chi/v5/middleware"
	_"keeper/docs" 
	httpSwagger "github.com/swaggo/http-swagger"
)

// APIServer represents the HTTP API server with its dependencies
type APIServer struct {
	listenAddr string             // Server listen address
	store      storage.Store      // Data storage interface
	validate   *validator.Validate // Request validation instance
	Router     *chi.Mux          // HTTP router instance
}

// NewAPIServer creates a new API server instance with configured routes and middleware
func NewAPIServer(listenAddr string, store storage.Store, validate *validator.Validate) *APIServer {
	server := &APIServer{
		listenAddr: listenAddr,
		store:      store,
		validate:   validate, 
		Router:     chi.NewRouter(),
	}

	// Configure middleware stack
	server.Router.Use(middleware.Logger)    // Request logging
	server.Router.Use(middleware.Recoverer) // Panic recovery

	// Health check endpoint
	server.Router.Get("/healthcheck", server.handleHealthCheck) // GET /healthcheck
	
	// API documentation endpoint
	server.Router.Get("/swagger/*", httpSwagger.WrapHandler)

	// Dealership resource routes
	server.Router.Route("/dealerships", func(r chi.Router) {
		r.Post("/", server.handleCreateDealership)     // Create new dealership
		r.Get("/", server.handleGetDealerships)        // List all dealerships
		r.Put("/{id}", server.handleUpdateDealership)  // Update existing dealership
		r.Delete("/{id}", server.handleDeleteDealership) // Delete dealership
	})

	// Employee resource routes
	server.Router.Route("/employees", func(r chi.Router) {
		r.Post("/", server.handleCreateEmployee)     // Create new employee
		r.Get("/", server.handleGetEmployee)         // List all employees
		r.Put("/{id}", server.handleUpdateEmployee)  // Update existing employee
		r.Delete("/{id}", server.handleDeleteEmployee) // Delete employee
	})

	// Employment resource routes
	server.Router.Route("/employments", func(r chi.Router) {
		r.Post("/", server.handleCreateEmployment)     // Create new employment
		r.Get("/", server.handleGetEmployments)        // List all employments
		r.Put("/{id}", server.handleUpdateEmployment)  // Update existing employment
		r.Delete("/{id}", server.handleDeleteEmployment) // Delete employment
	})

	// Client resource routes
	server.Router.Route("/clients", func(r chi.Router) {
		r.Post("/", server.handleCreateClient)     // Create new client
		r.Get("/", server.handleGetClients)        // List all clients
		r.Put("/{id}", server.handleUpdateClient)  // Update existing client
		r.Delete("/{id}", server.handleDeleteClient) // Delete client
	})

	// Car resource routes
	server.Router.Route("/car", func(r chi.Router) {
		r.Post("/", server.handleCreateCar)      // Create new car
		r.Get("/", server.handleGetCars)         // List all cars
		r.Patch("/{id}", server.handlePatchCar) // Partially update car
		r.Delete("/{id}", server.handleDeleteCar) // Delete car
	})

	// Order resource routes
	server.Router.Route("/orders", func(r chi.Router) {
		r.Post("/", server.handleCreateOrder)     // Create new order
		r.Get("/", server.handleGetOrders)        // List all orders
		r.Put("/{id}", server.handleUpdateOrder)  // Update existing order
		r.Delete("/{id}", server.handleDeleteOrder) // Delete order
	})

	// Appointment resource routes
	server.Router.Route("/appointments", func(r chi.Router) {
		r.Post("/", server.handleCreateAppointment)     // Create new appointment
		r.Get("/", server.handleGetAppointments)         // List all appointments
		r.Put("/{id}", server.handleUpdateAppointment)  // Update existing appointment
		r.Delete("/{id}", server.handleDeleteAppointment) // Delete appointment
	})
	
	return server
} 

// Run starts the HTTP server on the configured address
func (s *APIServer) Run() {
	log.Println("JSON API server running on port", s.listenAddr)
	http.ListenAndServe(s.listenAddr, s.Router)
}

// validateRequest validates incoming request data using the validator instance
// Returns true if validation passes, false otherwise (also writes error response)
func (s *APIServer) validateRequest(w http.ResponseWriter, r *http.Request, data any) bool {
	if err := s.validate.Struct(data); err != nil {
		logError(r, err)
		writeError(w, http.StatusBadRequest, err)
		return false
	}
	return true
}