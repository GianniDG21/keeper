// internal/api/server.go
package api

import (
	"keeper/internal/storage"
	"log"
	"net/http"
	"github.com/go-playground/validator/v10"
	"github.com/go-chi/chi/v5"          
	"github.com/go-chi/chi/v5/middleware" 
)

type APIServer struct {
	listenAddr string
	store      storage.Store
	validate  *validator.Validate
}

func NewAPIServer(listenAddr string, store storage.Store, validate *validator.Validate) *APIServer {
    return &APIServer{
        listenAddr: listenAddr,
        store:      store,
        validate:   validate, 
    }
}

func (s *APIServer) Run() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)    
	router.Use(middleware.Recoverer) 


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

	router.Route("/car", func(r chi.Router) {
		r.Post("/", s.handleCreateCar)   // POST /car
		r.Get("/", s.handleGetCars)	  // GET /car
		r.Put("/{id}", s.handleUpdateCar)  // PUT /car/{id}
		r.Delete("/{id}", s.handleDeleteCar) // DELETE /car/{id}
	})

	router.Route("/orders", func(r chi.Router) {
		r.Post("/", s.handleCreateOrder)   // POST /orders
		r.Get("/", s.handleGetOrders)	  // GET /orders
		r.Put("/{id}", s.handleUpdateOrder)  // PUT /orders/{id}
		r.Delete("/{id}", s.handleDeleteOrder) // DELETE /orders/{id}
	})

	router.Route("/appointments", func(r chi.Router) {
		r.Post("/", s.handleCreateAppointment)   // POST /appointments
		r.Get("/", s.handleGetAppointments)	  // GET /appointments
		r.Put("/{id}", s.handleUpdateAppointment)  // PUT /appointments/{id}
		r.Delete("/{id}", s.handleDeleteAppointment) // DELETE /appointments/{id}
	})

	log.Println("JSON API server running on port", s.listenAddr)

	
	err := http.ListenAndServe(s.listenAddr, router)
	if err != nil {
		log.Fatal(err)
	}
}
func (s *APIServer) validateRequest(w http.ResponseWriter, r *http.Request, data any) bool {
	if err := s.validate.Struct(data); err != nil {
		logError(r, err)
		writeError(w, http.StatusBadRequest, err)
		return false
	}
	return true
}