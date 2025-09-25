package api

import (
	"encoding/json"
	"errors"
	"keeper/internal/models"
	"net/http"
	"strconv"
	"github.com/go-chi/chi/v5"
)

// utils.writeJSON and utils.writeError are utility functions for writing JSON responses and error messages to the HTTP response writer.

// Function to check the server status
func (s *APIServer) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
    writeJSON(w, http.StatusOK, map[string]string{"status": "available"})
}

// Dealerships Handlers //
func (s *APIServer) handleCreateDealership(w http.ResponseWriter, r *http.Request) {
	var newDealership models.Dealership
	if err := json.NewDecoder(r.Body).Decode(&newDealership); err != nil {
		writeError(w, http.StatusBadRequest, err)
		logError(r, err)
		return
	}

	if !s.validateRequest(w, r, &newDealership) {
		return
	}

	newID, err := s.store.CreateDealership(&newDealership)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}
	writeJSON(w, http.StatusCreated, map[string]int{"id": newID})
}

func (s *APIServer) handleGetDealerships(w http.ResponseWriter, r *http.Request) {
	dealerships, err := s.store.GetDealerships()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}

	writeJSON(w, http.StatusOK, dealerships)
}

func (s *APIServer) handleUpdateDealership(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, errors.New("invalid ID format"))
		logError(r, err)
		return
	}

	var updatedDealership models.Dealership
	if err := json.NewDecoder(r.Body).Decode(&updatedDealership); err != nil {
		writeError(w, http.StatusBadRequest, err)
		logError(r, err)
		return
	}

	if !s.validateRequest(w, r, &updatedDealership) {
		return
	}

	if err := s.store.UpdateDealership(id, &updatedDealership); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}
	writeJSON(w, http.StatusOK, updatedDealership)
}

func (s *APIServer) handleDeleteDealership(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, errors.New("invalid ID format"))
		logError(r, err)
		return
	}

	if err := s.store.DeleteDealership(id); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}
	writeJSON(w, http.StatusNoContent, nil)
}

// Employees Handlers //
func (s *APIServer) handleCreateEmployee(w http.ResponseWriter, r *http.Request) {
	var newEmployee models.Employee
	if err := json.NewDecoder(r.Body).Decode(&newEmployee); err != nil {
		writeError(w, http.StatusBadRequest, err)
		logError(r, err)
		return
	}

	if !s.validateRequest(w, r, &newEmployee) {
		return
	}

	newID, err := s.store.CreateEmployee(&newEmployee)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}
	writeJSON(w, http.StatusCreated, map[string]int{"id": newID})
}

func (s *APIServer) handleGetEmployee(w http.ResponseWriter, r *http.Request) {
	employees, err := s.store.GetEmployee()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}

	writeJSON(w, http.StatusOK, employees)
}

func (s *APIServer) handleUpdateEmployee(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, errors.New("invalid ID format"))
		logError(r, err)
		return
	}

	var updatedEmployee models.Employee
	if err := json.NewDecoder(r.Body).Decode(&updatedEmployee); err != nil {
		writeError(w, http.StatusBadRequest, err)
		logError(r, err)
		return
	}

	if !s.validateRequest(w, r, &updatedEmployee) {
		return
	}

	if err := s.store.UpdateEmployee(id, &updatedEmployee); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}
	writeJSON(w, http.StatusOK, updatedEmployee)
}

func (s *APIServer) handleDeleteEmployee(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, errors.New("invalid ID format"))
		logError(r, err)
		return
	}

	if err := s.store.DeleteEmployee(id); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}
	writeJSON(w, http.StatusNoContent, nil)
}

// Employments Handlers //
func (s *APIServer) handleCreateEmployment(w http.ResponseWriter, r *http.Request) {
	var newEmployment models.Employment
	if err := json.NewDecoder(r.Body).Decode(&newEmployment); err != nil {
		writeError(w, http.StatusBadRequest, err)
		logError(r, err)
		return
	}

	if !s.validateRequest(w, r, &newEmployment) {
		return
	}

	newID, err := s.store.CreateEmployment(&newEmployment)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}
	writeJSON(w, http.StatusCreated, map[string]int{"id": newID})
}

func (s *APIServer) handleGetEmployments(w http.ResponseWriter, r *http.Request) {
	employments, err := s.store.GetEmployments()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}

	writeJSON(w, http.StatusOK, employments)
}

func (s *APIServer) handleUpdateEmployment(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, errors.New("invalid ID format"))
		logError(r, err)
		return
	}

	var updatedEmployment models.Employment
	if err := json.NewDecoder(r.Body).Decode(&updatedEmployment); err != nil {
		writeError(w, http.StatusBadRequest, err)
		logError(r, err)
		return
	}

	if !s.validateRequest(w, r, &updatedEmployment) {
		return
	}

	if err := s.store.UpdateEmployment(id, &updatedEmployment); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}
	writeJSON(w, http.StatusOK, updatedEmployment)
}

func (s *APIServer) handleDeleteEmployment(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, errors.New("invalid ID format"))
		logError(r, err)
		return
	}

	if err := s.store.DeleteEmployment(id); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}
	writeJSON(w, http.StatusNoContent, nil)
}

// Clients Handlers //

func (s *APIServer) handleCreateClient(w http.ResponseWriter, r *http.Request) {
	var newClient models.Client
	if err := json.NewDecoder(r.Body).Decode(&newClient); err != nil {
		writeError(w, http.StatusBadRequest, err)
		logError(r, err)
		return
	}
	
	if !s.validateRequest(w, r, &newClient) {
		return
	}

	newID, err := s.store.CreateClient(&newClient)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}
	writeJSON(w, http.StatusCreated, map[string]int{"id": newID})
}

func (s *APIServer) handleGetClients(w http.ResponseWriter, r *http.Request) {
	clients, err := s.store.GetClients()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}

	writeJSON(w, http.StatusOK, clients)
}

func (s *APIServer) handleUpdateClient(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, errors.New("invalid ID format"))
		logError(r, err)
		return
	}

	var updatedClient models.Client
	if err := json.NewDecoder(r.Body).Decode(&updatedClient); err != nil {
		writeError(w, http.StatusBadRequest, err)
		logError(r, err)
		return
	}

	if !s.validateRequest(w, r, &updatedClient) {
		return
	}

	if err := s.store.UpdateClient(id, &updatedClient); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}
	writeJSON(w, http.StatusOK, updatedClient)
}

func (s *APIServer) handleDeleteClient(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, errors.New("invalid ID format"))
		logError(r, err)
		return
	}

	if err := s.store.DeleteClient(id); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}
	writeJSON(w, http.StatusNoContent, nil)
}

// CarPark Handlers //
func (s *APIServer) handleCreateCar(w http.ResponseWriter, r *http.Request) {
	var newCar models.CarPark
	if err := json.NewDecoder(r.Body).Decode(&newCar); err != nil {
		writeError(w, http.StatusBadRequest, err)
		logError(r, err)
		return
	}

	if !s.validateRequest(w, r, &newCar) {
		return
	}

	newID, err := s.store.CreateCarPark(&newCar)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}
	writeJSON(w, http.StatusCreated, map[string]int{"id": newID})
}

func (s *APIServer) handleGetCars(w http.ResponseWriter, r *http.Request) {
	cars, err := s.store.GetCarParks()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}
	writeJSON(w, http.StatusOK, cars)
}

func (s *APIServer) handleUpdateCar(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, errors.New("invalid ID format"))
		logError(r, err)
		return
	}

	var updatedCar models.CarPark
	if err := json.NewDecoder(r.Body).Decode(&updatedCar); err != nil {
		writeError(w, http.StatusBadRequest, err)
		logError(r, err)
		return
	}

	if !s.validateRequest(w, r, &updatedCar) {
		return
	}

	if err := s.store.UpdateCarPark(id, &updatedCar); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}
	writeJSON(w, http.StatusOK, updatedCar)
}

func (s *APIServer) handleDeleteCar(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, errors.New("invalid ID format"))
		logError(r, err)
		return
	}

	if err := s.store.DeleteCarPark(id); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}
	writeJSON(w, http.StatusNoContent, nil)
}


// Orders Handlers //
func (s *APIServer) handleCreateOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder models.Order
	if err := json.NewDecoder(r.Body).Decode(&newOrder); err != nil {
		writeError(w, http.StatusBadRequest, err)
		logError(r, err)
		return
	}

	if !s.validateRequest(w, r, &newOrder) {
		return
	}

	newID, err := s.store.CreateOrder(&newOrder)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}
	writeJSON(w, http.StatusCreated, map[string]int{"id": newID})
}

func (s *APIServer) handleGetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := s.store.GetOrders()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}
	writeJSON(w, http.StatusOK, orders)
}

func (s *APIServer) handleUpdateOrder(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, errors.New("invalid ID format"))
		logError(r, err)
		return
	}

	var updatedOrder models.Order
	if err := json.NewDecoder(r.Body).Decode(&updatedOrder); err != nil {
		writeError(w, http.StatusBadRequest, err)
		logError(r, err)
		return
	}

	if !s.validateRequest(w, r, &updatedOrder) {
		return
	}

	if err := s.store.UpdateOrder(id, &updatedOrder); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}
	writeJSON(w, http.StatusOK, updatedOrder)
}

func (s *APIServer) handleDeleteOrder(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, errors.New("invalid ID format"))
		logError(r, err)
		return
	}

	if err := s.store.DeleteOrder(id); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}
	writeJSON(w, http.StatusNoContent, nil)
}

// Appointments Handlers //
func (s *APIServer) handleCreateAppointment(w http.ResponseWriter, r *http.Request) {
	var newAppointment models.Appointment
	if err := json.NewDecoder(r.Body).Decode(&newAppointment); err != nil {
		writeError(w, http.StatusBadRequest, err)
		logError(r, err)
		return
	}

	if !s.validateRequest(w, r, &newAppointment) {
		return
	}

	newID, err := s.store.CreateAppointment(&newAppointment)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}
	writeJSON(w, http.StatusCreated, map[string]int{"id": newID})
}

func (s *APIServer) handleGetAppointments(w http.ResponseWriter, r *http.Request) {
	appointments, err := s.store.GetAppointments()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}
	writeJSON(w, http.StatusOK, appointments)
}

func (s *APIServer) handleUpdateAppointment(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, errors.New("invalid ID format"))
		logError(r, err)
		return
	}

	var updatedAppointment models.Appointment
	if err := json.NewDecoder(r.Body).Decode(&updatedAppointment); err != nil {
		writeError(w, http.StatusBadRequest, err)
		logError(r, err)
		return
	}

	if !s.validateRequest(w, r, &updatedAppointment) {
		return
	}

	if err := s.store.UpdateAppointment(id, &updatedAppointment); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}
	writeJSON(w, http.StatusOK, updatedAppointment)
}

func (s *APIServer) handleDeleteAppointment(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, errors.New("invalid ID format"))
		logError(r, err)
		return
	}

	if err := s.store.DeleteAppointment(id); err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}
	writeJSON(w, http.StatusNoContent, nil)
}

