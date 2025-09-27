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

type APIServer struct {
	listenAddr string
	store      storage.Store
	validate  *validator.Validate
	Router     *chi.Mux
}

func NewAPIServer(listenAddr string, store storage.Store, validate *validator.Validate) *APIServer {
    server := &APIServer{
        listenAddr: listenAddr,
        store:      store,
        validate:   validate, 
		Router:     chi.NewRouter(),
    }

	server.Router.Use(middleware.Logger)    
	server.Router.Use(middleware.Recoverer) 


	server.Router.Get("/healthcheck", server.handleHealthCheck) // GET /healthcheck
	server.Router.Get("/swagger/*", httpSwagger.WrapHandler)

	server.Router.Route("/dealerships", func(r chi.Router) {
		r.Post("/", server.handleCreateDealership)   // POST /dealerships
		r.Get("/", server.handleGetDealerships)      // GET /dealerships
		r.Put("/{id}", server.handleUpdateDealership)  // PUT /dealerships/{id}
		r.Delete("/{id}", server.handleDeleteDealership) // DELETE /dealerships/{id}
	})

	server.Router.Route("/employees", func(r chi.Router) {
		r.Post("/", server.handleCreateEmployee)   // POST /employees
		r.Get("/", server.handleGetEmployee)      // GET /employees
		r.Put("/{id}", server.handleUpdateEmployee)  // PUT /employees/{id}
		r.Delete("/{id}", server.handleDeleteEmployee) // DELETE /employees/{id}
	})

	server.Router.Route("/employments", func(r chi.Router) {
		r.Post("/", server.handleCreateEmployment)   // POST /employments
		r.Get("/", server.handleGetEmployments)      // GET /employments
		r.Put("/{id}", server.handleUpdateEmployment)  // PUT /employments/{id}
		r.Delete("/{id}", server.handleDeleteEmployment) // DELETE /employments/{id}
	})

	server.Router.Route("/clients", func(r chi.Router) {
		r.Post("/", server.handleCreateClient)   // POST /clients
		r.Get("/", server.handleGetClients)      // GET /clients
		r.Put("/{id}", server.handleUpdateClient)  // PUT /clients/{id}
		r.Delete("/{id}", server.handleDeleteClient) // DELETE /clients/{id}
	})

	server.Router.Route("/car", func(r chi.Router) {
		r.Post("/", server.handleCreateCar)   // POST /car
		r.Get("/", server.handleGetCars)	  // GET /car
		r.Put("/{id}", server.handleUpdateCar)  // PUT /car/{id}
		r.Delete("/{id}", server.handleDeleteCar) // DELETE /car/{id}
	})

	server.Router.Route("/orders", func(r chi.Router) {
		r.Post("/", server.handleCreateOrder)   // POST /orders
		r.Get("/", server.handleGetOrders)	  // GET /orders
		r.Put("/{id}", server.handleUpdateOrder)  // PUT /orders/{id}
		r.Delete("/{id}", server.handleDeleteOrder) // DELETE /orders/{id}
	})

	server.Router.Route("/appointments", func(r chi.Router) {
		r.Post("/", server.handleCreateAppointment)   // POST /appointments
		r.Get("/", server.handleGetAppointments)	  // GET /appointments
		r.Put("/{id}", server.handleUpdateAppointment)  // PUT /appointments/{id}
		r.Delete("/{id}", server.handleDeleteAppointment) // DELETE /appointments/{id}
	})
	return server
} 

func (s *APIServer) Run() {
	log.Println("JSON API server running on port", s.listenAddr)
	http.ListenAndServe(s.listenAddr, s.Router)

}
func (s *APIServer) validateRequest(w http.ResponseWriter, r *http.Request, data any) bool {
	if err := s.validate.Struct(data); err != nil {
		logError(r, err)
		writeError(w, http.StatusBadRequest, err)
		return false
	}
	return true
}