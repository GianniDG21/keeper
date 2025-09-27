// in internal/api/handlers_test.go
package api

import (
	"bytes"
	"encoding/json"
	"keeper/internal/models"
	"keeper/internal/storage"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-playground/validator/v10"
)

func newTestServer(_ *testing.T, store storage.Store) *APIServer {
	validate := validator.New()
	return NewAPIServer(":0", store, validate)
}

func newTestDB(t *testing.T) *storage.PostgresStore {
	connString := os.Getenv("TEST_DATABASE_URL")
	if connString == "" {
		t.Skip("TEST_DATABASE_URL not set, skipping integration test")
	}

	store, err := storage.NewPostgresStore(connString)
	if err != nil {
		t.Fatalf("failed to connect to test database: %s", err)
	}

	_, err = store.Db.Exec(`TRUNCATE TABLE dealership, employee, employment, car_park, client, appointment, "order" RESTART IDENTITY CASCADE;`)
	if err != nil {
		t.Fatalf("failed to clean test database: %s", err)
	}

	return store
}

func TestCreateDealershipAPI(t *testing.T) {
	store := newTestDB(t)
	
	server := newTestServer(t, store)

	testServer := httptest.NewServer(server.Router) 
	defer testServer.Close()

	t.Run("it creates a dealership and returns its ID", func(t *testing.T) {
		payload := []byte(`{"postal_code":"73100","city":"Test City","address":"Test Address","phone":"5551234"}`)
		reqBody := bytes.NewBuffer(payload)

		resp, err := http.Post(testServer.URL+"/dealerships", "application/json", reqBody)
		if err != nil {
			t.Fatalf("failed to send request: %s", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			t.Errorf("expected status %d, got %d", http.StatusCreated, resp.StatusCode)
		}

		var result map[string]int
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			t.Fatalf("failed to decode response body: %s", err)
		}
		
		id, ok := result["id"]
		if !ok || id == 0 {
			t.Errorf("response does not contain a valid new ID: %v", result)
		}
	})
}

func TestGetDealershipsAPI(t *testing.T) {
	
	store := newTestDB(t)
	server := newTestServer(t, store)
	testServer := httptest.NewServer(server.Router)
	defer testServer.Close()

	// SEEDING
	_, err := store.Db.Exec("INSERT INTO dealership (postalcode, city, address, phone) VALUES ('73100', 'Lecce Seed', 'Via Seed 1', '12345')")
	if err != nil {
		t.Fatalf("Impossibile inserire dati di prova: %v", err)
	}

	resp, err := http.Get(testServer.URL + "/dealerships")
	if err != nil {
		t.Fatalf("Errore durante l'invio della richiesta: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status code errato: ricevuto %v, atteso %v", resp.StatusCode, http.StatusOK)
	}
	
	var dealerships []*models.Dealership
	if err := json.NewDecoder(resp.Body).Decode(&dealerships); err != nil {
		t.Fatalf("impossibile decodificare la risposta JSON: %s", err)
	}

	if len(dealerships) != 1 {
		t.Errorf("attesa 1 concessionaria, ricevute %d", len(dealerships))
	}

	if dealerships[0].City != "Lecce Seed" {
		t.Errorf("città errata: ricevuta '%s', attesa 'Lecce Seed'", dealerships[0].City)
	}
}