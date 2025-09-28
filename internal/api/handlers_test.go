// Package api provides HTTP handlers and API server functionality for the KEEPER application.
// This file contains integration tests for the API endpoints.
package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"keeper/internal/models"
	"keeper/internal/storage"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-playground/validator/v10"
)

// errStatusMismatch defines a standard error message format for status code mismatches in tests.
const errStatusMismatch = "wrong status code: received %v, expected %v"

// newTestServer creates a new APIServer instance for testing purposes.
// It accepts a testing instance and a storage store, returning a configured server.
func newTestServer(_ *testing.T, store storage.Store) *APIServer {
	validate := validator.New()
	return NewAPIServer(":0", store, validate)
}

// newTestDB establishes a connection to the test database and performs cleanup.
// It reads the connection string from TEST_DATABASE_URL environment variable.
// Returns a PostgresStore instance ready for testing.
func newTestDB(t *testing.T) *storage.PostgresStore {
	connString := os.Getenv("TEST_DATABASE_URL")
	if connString == "" {
		t.Skip("TEST_DATABASE_URL not set, skipping integration test")
	}

	store, err := storage.NewPostgresStore(connString)
	if err != nil {
		t.Fatalf("failed to connect to test database: %s", err)
	}

	// Clean all tables and reset identity sequences to ensure test isolation
	_, err = store.Db.Exec(`TRUNCATE TABLE dealership, employee, employment, car_park, client, appointment, "order" RESTART IDENTITY CASCADE;`)
	if err != nil {
		t.Fatalf("failed to clean test database: %s", err)
	}

	return store
}

// TestCreateDealershipAPI tests the POST /dealerships endpoint.
// Verifies that a dealership can be created successfully and returns the correct response.
func TestCreateDealershipAPI(t *testing.T) {
	store := newTestDB(t)
	
	server := newTestServer(t, store)

	testServer := httptest.NewServer(server.Router) 
	defer testServer.Close()

	t.Run("it creates a dealership and returns its ID", func(t *testing.T) {
		// Prepare test payload with valid dealership data
		payload := []byte(`{"postal_code":"73100","city":"Test City","address":"Test Address","phone":"5551234"}`)
		reqBody := bytes.NewBuffer(payload)

		// Send POST request to create dealership
		resp, err := http.Post(testServer.URL+"/dealerships", "application/json", reqBody)
		if err != nil {
			t.Fatalf("failed to send request: %s", err)
		}
		defer resp.Body.Close()

		// Verify response status code
		if resp.StatusCode != http.StatusCreated {
			t.Errorf(errStatusMismatch, resp.StatusCode, http.StatusCreated)
		}

		// Parse response body and verify dealership ID is returned
		var result map[string]int
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			t.Fatalf("failed to decode response body: %s", err)
		}
		
		id, ok := result["id"]
		if !ok || id == 0 {
			t.Errorf("expected valid dealership ID in response, got: %v", result)
		}
	})
}

// TestGetDealershipsAPI tests the GET /dealerships endpoint.
// Verifies that dealerships can be retrieved successfully and returns the correct data.
func TestGetDealershipsAPI(t *testing.T) {
	
	store := newTestDB(t)
	server := newTestServer(t, store)
	testServer := httptest.NewServer(server.Router)
	defer testServer.Close()

	// Insert test data into the database
	_, err := store.Db.Exec("INSERT INTO dealership (postalcode, city, address, phone) VALUES ('73100', 'Lecce Seed', 'Via Seed 1', '12345')")
	if err != nil {
		t.Fatalf("unable to insert test data: %v", err)
	}

	// Send GET request to retrieve dealerships
	resp, err := http.Get(testServer.URL + "/dealerships")
	if err != nil {
		t.Fatalf("error during request sending: %s", err)
	}
	defer resp.Body.Close()

	// Verify response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf(errStatusMismatch, resp.StatusCode, http.StatusOK)
	}
	
	// Parse response body and verify dealership data
	var dealerships []*models.Dealership
	if err := json.NewDecoder(resp.Body).Decode(&dealerships); err != nil {
		t.Fatalf("unable to decode JSON response: %s", err)
	}

	// Verify that exactly one dealership is returned
	if len(dealerships) != 1 {
		t.Errorf("expected 1 dealership, received %d", len(dealerships))
	}

	// Verify dealership data matches test data
	if dealerships[0].City != "Lecce Seed" {
		t.Errorf("wrong city: received '%s', expected 'Lecce Seed'", dealerships[0].City)
	}
}

// TestPatchVehicleAPI tests the PATCH /car/{id} endpoint.
// Verifies that a vehicle's data can be partially updated successfully.
func TestPatchVehicleAPI(t *testing.T) {
	store := newTestDB(t)
	server := newTestServer(t, store)
	testServer := httptest.NewServer(server.Router)
	defer testServer.Close()

	// Insert test dealership
	var dealershipID int
	err := store.Db.QueryRow("INSERT INTO dealership (postalcode, city, address, phone) VALUES ('00000', 'Test', 'Test', '123') RETURNING id_dealership").Scan(&dealershipID)
	if err != nil {
		t.Fatalf("unable to insert test dealership: %v", err)
	}

	// Insert test vehicle associated with the dealership
	var vehicleID int
	err = store.Db.QueryRow("INSERT INTO car_park (vin, id_dealership, brand, model, \"year\", km) VALUES ('TESTVINUPDATE0001', $1, 'Fiat', 'Panda', 2020, 50000) RETURNING id_car", dealershipID).Scan(&vehicleID)
	if err != nil {
		t.Fatalf("unable to insert test vehicle: %v", err)
	}

	// Prepare PATCH payload to update vehicle kilometers
	payload := []byte(`{"km": 60000}`)
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/car/%d", testServer.URL, vehicleID), bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("unable to create PATCH request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	
	// Send PATCH request to update vehicle
	resp, err := testServer.Client().Do(req)
	if err != nil {
		t.Fatalf("error during PATCH request sending: %s", err)
	}
	defer resp.Body.Close()

	// Verify response status code (Note: should be StatusOK, not StatusCreated)
	if resp.StatusCode != http.StatusOK {
		t.Errorf(errStatusMismatch, resp.StatusCode, http.StatusOK)
	}
}

// TestDeleteEmployeeAPI tests the DELETE /employees/{id} endpoint.
// Verifies that an employee can be deleted successfully and is removed from the database.
func TestDeleteEmployeeAPI(t *testing.T) {
	store := newTestDB(t)
	server := newTestServer(t, store)
	testServer := httptest.NewServer(server.Router)
	defer testServer.Close()

	// Insert test employee
	var employeeID int
	err := store.Db.QueryRow("INSERT INTO employee (tin, name, surname, role) VALUES ('TESTTINDELETE01', 'Marco', 'Verdi', 'salesperson') RETURNING id_employee").Scan(&employeeID)
	if err != nil {
		t.Fatalf("unable to insert test employee: %v", err)
	}

	// Create DELETE request for the employee
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/employees/%d", testServer.URL, employeeID), nil)
	if err != nil {
		t.Fatalf("unable to create DELETE request: %v", err)
	}
	
	// Send DELETE request
	resp, err := testServer.Client().Do(req)
	if err != nil {
		t.Fatalf("error during DELETE request sending: %s", err)
	}
	defer resp.Body.Close()

	// Verify response status code
	if resp.StatusCode != http.StatusNoContent {
		t.Errorf(errStatusMismatch, resp.StatusCode, http.StatusNoContent)
	}

	// Verify employee was actually deleted from the database
	var count int
	err = store.Db.QueryRow("SELECT COUNT(*) FROM employee WHERE id_employee = $1", employeeID).Scan(&count)
	if err != nil {
		t.Fatalf("error checking DB after deletion: %s", err)
	}

	if count != 0 {
		t.Errorf("employee record still exists in DB, but should have been deleted")
	}
}