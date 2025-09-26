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

// @Summary      Health Check
// @Description  Checks if the API server is running.
// @Tags         Health
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /healthcheck [get]
func (s *APIServer) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "available"})
}

// Dealerships Handlers //

// @Summary      Create a new Dealership
// @Description  Adds a new dealership to the system from the provided JSON payload.
// @Tags         Dealerships
// @Accept       json
// @Produce      json
// @Param        dealership  body      models.Dealership  true  "New Dealership Data"
// @Success      201         {object}  map[string]int     "Returns the ID of the newly created dealership"
// @Failure      400         {object}  map[string]string  "Error: Invalid request payload"
// @Failure      500         {object}  map[string]string  "Error: Internal server error"
// @Router       /dealerships [post]
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

// @Summary      List all Dealerships
// @Description  Retrieves a list of all dealership branches.
// @Tags         Dealerships
// @Produce      json
// @Success      200  {array}   models.Dealership
// @Failure      500  {object}  map[string]string "Error: Internal server error"
// @Router       /dealerships [get]
func (s *APIServer) handleGetDealerships(w http.ResponseWriter, r *http.Request) {
	dealerships, err := s.store.GetDealerships()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}

	writeJSON(w, http.StatusOK, dealerships)
}

// @Summary      Update a Dealership
// @Description  Updates an existing dealership's data by its ID.
// @Tags         Dealerships
// @Accept       json
// @Produce      json
// @Param        id          path      int                true  "Dealership ID"
// @Param        dealership  body      models.Dealership  true  "Updated Dealership Data"
// @Success      200         {object}  models.Dealership
// @Failure      400         {object}  map[string]string "Error: Invalid ID or request payload"
// @Failure      500         {object}  map[string]string "Error: Internal server error"
// @Router       /dealerships/{id} [put]
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

// @Summary      Delete a Dealership
// @Description  Deletes a dealership by its ID.
// @Tags         Dealerships
// @Produce      json
// @Param        id  path      int  true  "Dealership ID"
// @Success      204 "No Content"
// @Failure      400 {object}  map[string]string "Error: Invalid ID"
// @Failure      500 {object}  map[string]string "Error: Internal server error"
// @Router       /dealerships/{id} [delete]
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

// @Summary      Create Employee
// @Tags         Employees
// @Accept       json
// @Produce      json
// @Param        employee  body      models.Employee      true  "New Employee"
// @Success      201       {object}  map[string]int
// @Failure      400       {object}  map[string]string
// @Failure      500       {object}  map[string]string
// @Router       /employees [post]
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

// @Summary      List Employees
// @Tags         Employees
// @Produce      json
// @Success      200  {array}   models.Employee
// @Failure      500  {object}  map[string]string
// @Router       /employees [get]
func (s *APIServer) handleGetEmployee(w http.ResponseWriter, r *http.Request) {
	employees, err := s.store.GetEmployee()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}

	writeJSON(w, http.StatusOK, employees)
}

// @Summary      Update Employee
// @Tags         Employees
// @Accept       json
// @Produce      json
// @Param        id        path      int              true  "Employee ID"
// @Param        employee  body      models.Employee  true  "Updated Employee Data"
// @Success      200       {object}  models.Employee
// @Failure      400       {object}  map[string]string
// @Failure      500       {object}  map[string]string
// @Router       /employees/{id} [put]
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

// @Summary      Delete Employee
// @Tags         Employees
// @Produce      json
// @Param        id  path      int  true  "Employee ID"
// @Success      204 "No Content"
// @Failure      400 {object}  map[string]string
// @Failure      500 {object}  map[string]string
// @Router       /employees/{id} [delete]
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

// @Summary      Create Employment
// @Description  Assigns an employee to a dealership, creating a new employment record.
// @Tags         Employments
// @Accept       json
// @Produce      json
// @Param        employment  body      models.Employment    true  "New Employment Data"
// @Success      201         {object}  map[string]int     "Returns the ID of the new employment record"
// @Failure      400         {object}  map[string]string  "Error: Invalid request payload"
// @Failure      500         {object}  map[string]string  "Error: Internal server error"
// @Router       /employments [post]
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

// @Summary      List Employments
// @Description  Retrieves a list of all employment records.
// @Tags         Employments
// @Produce      json
// @Success      200  {array}   models.Employment
// @Failure      500  {object}  map[string]string "Error: Internal server error"
// @Router       /employments [get]
func (s *APIServer) handleGetEmployments(w http.ResponseWriter, r *http.Request) {
	employments, err := s.store.GetEmployments()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}

	writeJSON(w, http.StatusOK, employments)
}

// @Summary      Update Employment
// @Description  Updates an existing employment record by its ID (e.g., to set an end date).
// @Tags         Employments
// @Accept       json
// @Produce      json
// @Param        id          path      int                  true  "Employment ID"
// @Param        employment  body      models.Employment    true  "Updated Employment Data"
// @Success      200         {object}  models.Employment
// @Failure      400         {object}  map[string]string "Error: Invalid ID or request payload"
// @Failure      500         {object}  map[string]string "Error: Internal server error"
// @Router       /employments/{id} [put]
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

// @Summary      Delete Employment
// @Description  Deletes an employment record by its ID.
// @Tags         Employments
// @Produce      json
// @Param        id  path      int  true  "Employment ID"
// @Success      204 "No Content"
// @Failure      400 {object}  map[string]string "Error: Invalid ID"
// @Failure      500 {object}  map[string]string "Error: Internal server error"
// @Router       /employments/{id} [delete]
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

// @Summary      Create Client
// @Description  Registers a new client (private or business) in the system.
// @Tags         Clients
// @Accept       json
// @Produce      json
// @Param        client  body      models.Client        true  "New Client Data"
// @Success      201     {object}  map[string]int     "Returns the ID of the newly created client"
// @Failure      400     {object}  map[string]string  "Error: Invalid request payload"
// @Failure      500     {object}  map[string]string  "Error: Internal server error"
// @Router       /clients [post]
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

// @Summary      List Clients
// @Description  Retrieves a list of all clients.
// @Tags         Clients
// @Produce      json
// @Success      200  {array}   models.Client
// @Failure      500  {object}  map[string]string "Error: Internal server error"
// @Router       /clients [get]
func (s *APIServer) handleGetClients(w http.ResponseWriter, r *http.Request) {
	clients, err := s.store.GetClients()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}

	writeJSON(w, http.StatusOK, clients)
}

// @Summary      Update Client
// @Description  Updates an existing client's data by their ID.
// @Tags         Clients
// @Accept       json
// @Produce      json
// @Param        id      path      int              true  "Client ID"
// @Param        client  body      models.Client    true  "Updated Client Data"
// @Success      200     {object}  models.Client
// @Failure      400     {object}  map[string]string "Error: Invalid ID or request payload"
// @Failure      500     {object}  map[string]string "Error: Internal server error"
// @Router       /clients/{id} [put]
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

// @Summary      Delete Client
// @Description  Deletes a client by their ID.
// @Tags         Clients
// @Produce      json
// @Param        id  path      int  true  "Client ID"
// @Success      204 "No Content"
// @Failure      400 {object}  map[string]string "Error: Invalid ID"
// @Failure      500 {object}  map[string]string "Error: Internal server error"
// @Router       /clients/{id} [delete]
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

// @Summary      Add a new Vehicle
// @Description  Adds a new vehicle to the inventory.
// @Tags         Vehicles
// @Accept       json
// @Produce      json
// @Param        vehicle  body      models.CarPark       true  "New Vehicle Data"
// @Success      201      {object}  map[string]int     "Returns the ID of the newly created vehicle"
// @Failure      400      {object}  map[string]string  "Error: Invalid request payload"
// @Failure      500      {object}  map[string]string  "Error: Internal server error"
// @Router       /vehicles [post]
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

// @Summary      List all Vehicles
// @Description  Retrieves a list of all vehicles in the car park.
// @Tags         Vehicles
// @Produce      json
// @Success      200  {array}   models.CarPark
// @Failure      500  {object}  map[string]string "Error: Internal server error"
// @Router       /vehicles [get]
func (s *APIServer) handleGetCars(w http.ResponseWriter, r *http.Request) {
	cars, err := s.store.GetCarParks()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}
	writeJSON(w, http.StatusOK, cars)
}

// @Summary      Update a Vehicle
// @Description  Updates an existing vehicle's data by its ID.
// @Tags         Vehicles
// @Accept       json
// @Produce      json
// @Param        id       path      int              true  "Vehicle ID"
// @Param        vehicle  body      models.CarPark   true  "Updated Vehicle Data"
// @Success      200      {object}  models.CarPark
// @Failure      400      {object}  map[string]string "Error: Invalid ID or request payload"
// @Failure      500      {object}  map[string]string "Error: Internal server error"
// @Router       /vehicles/{id} [put]
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

// @Summary      Delete a Vehicle
// @Description  Deletes a vehicle from the inventory by its ID.
// @Tags         Vehicles
// @Produce      json
// @Param        id  path      int  true  "Vehicle ID"
// @Success      204 "No Content"
// @Failure      400 {object}  map[string]string "Error: Invalid ID"
// @Failure      500 {object}  map[string]string "Error: Internal server error"
// @Router       /vehicles/{id} [delete]
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

// @Summary      Create a new Order
// @Description  Creates a new sales order, linking a client, employee, and vehicle.
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param        order  body      models.Order         true  "New Order Data"
// @Success      201    {object}  map[string]int     "Returns the ID of the newly created order"
// @Failure      400    {object}  map[string]string  "Error: Invalid request payload"
// @Failure      500    {object}  map[string]string  "Error: Internal server error"
// @Router       /orders [post]
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

// @Summary      List all Orders
// @Description  Retrieves a list of all sales orders.
// @Tags         Orders
// @Produce      json
// @Success      200  {array}   models.Order
// @Failure      500  {object}  map[string]string "Error: Internal server error"
// @Router       /orders [get]
func (s *APIServer) handleGetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := s.store.GetOrders()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}
	writeJSON(w, http.StatusOK, orders)
}

// @Summary      Update an Order
// @Description  Updates an existing order's data (e.g., status) by its ID.
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param        id     path      int            true  "Order ID"
// @Param        order  body      models.Order   true  "Updated Order Data"
// @Success      200    {object}  models.Order
// @Failure      400    {object}  map[string]string "Error: Invalid ID or request payload"
// @Failure      500    {object}  map[string]string "Error: Internal server error"
// @Router       /orders/{id} [put]
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

// @Summary      Delete an Order
// @Description  Deletes a sales order by its ID.
// @Tags         Orders
// @Produce      json
// @Param        id  path      int  true  "Order ID"
// @Success      204 "No Content"
// @Failure      400 {object}  map[string]string "Error: Invalid ID"
// @Failure      500 {object}  map[string]string "Error: Internal server error"
// @Router       /orders/{id} [delete]
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

// @Summary      Create a new Appointment
// @Description  Schedules a new appointment (e.g., test drive, consultation).
// @Tags         Appointments
// @Accept       json
// @Produce      json
// @Param        appointment  body      models.Appointment   true  "New Appointment Data"
// @Success      201          {object}  map[string]int     "Returns the ID of the newly created appointment"
// @Failure      400          {object}  map[string]string  "Error: Invalid request payload"
// @Failure      500          {object}  map[string]string  "Error: Internal server error"
// @Router       /appointments [post]
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

// @Summary      List all Appointments
// @Description  Retrieves a list of all scheduled appointments.
// @Tags         Appointments
// @Produce      json
// @Success      200  {array}   models.Appointment
// @Failure      500  {object}  map[string]string "Error: Internal server error"
// @Router       /appointments [get]
func (s *APIServer) handleGetAppointments(w http.ResponseWriter, r *http.Request) {
	appointments, err := s.store.GetAppointments()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		logError(r, err)
		return
	}
	writeJSON(w, http.StatusOK, appointments)
}

// @Summary      Update an Appointment
// @Description  Updates an existing appointment by its ID (e.g., to reschedule).
// @Tags         Appointments
// @Accept       json
// @Produce      json
// @Param        id           path      int                  true  "Appointment ID"
// @Param        appointment  body      models.Appointment   true  "Updated Appointment Data"
// @Success      200          {object}  models.Appointment
// @Failure      400          {object}  map[string]string "Error: Invalid ID or request payload"
// @Failure      500          {object}  map[string]string "Error: Internal server error"
// @Router       /appointments/{id} [put]
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

// @Summary      Delete an Appointment
// @Description  Cancels and deletes an appointment by its ID.
// @Tags         Appointments
// @Produce      json
// @Param        id  path      int  true  "Appointment ID"
// @Success      204 "No Content"
// @Failure      400 {object}  map[string]string "Error: Invalid ID"
// @Failure      500 {object}  map[string]string "Error: Internal server error"
// @Router       /appointments/{id} [delete]
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
