//internal/api/utils.go

package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

// writeJSON writes a JSON response with the given status code and data
func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// writeError writes an error response in JSON format
func writeError(w http.ResponseWriter, status int, err error) {
	writeJSON(w, status, map[string]string{"error": err.Error()})
}

// logError logs HTTP request errors with method and path context
func logError(r *http.Request, err error) {
	log.Printf("[%s %s] ERROR: %v", r.Method, r.URL.Path, err)
}

// getIDFromURL extracts and validates an integer ID from the URL path parameter
func getIDFromURL(r *http.Request) (int, error) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New("invalid ID format")
	}
	return id, nil
}

func TestWriteJSON(t *testing.T) {
    tests := []struct {
        name       string
        statusCode int
        data       interface{}
        expected   string
    }{
        {
            name:       "success with map",
            statusCode: 200,
            data:       map[string]string{"key": "value"},
            expected:   `{"key":"value"}`,
        },
        {
            name:       "success with struct",
            statusCode: 201,
            data:       struct{Name string `json:"name"`}{Name: "test"},
            expected:   `{"name":"test"}`,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            w := httptest.NewRecorder()
            
            err := writeJSON(w, tt.statusCode, tt.data)
            
            if err != nil {
                t.Errorf("writeJSON() error = %v", err)
            }
            
            // Check status code
            if w.Code != tt.statusCode {
                t.Errorf("Expected status %d, got %d", tt.statusCode, w.Code)
            }
            
            // Check Content-Type header
            expectedContentType := "application/json"
            if ct := w.Header().Get("Content-Type"); ct != expectedContentType {
                t.Errorf("Expected Content-Type %s, got %s", expectedContentType, ct)
            }
            
            // Check JSON content
            body := strings.TrimSpace(w.Body.String())
            if body != tt.expected {
                t.Errorf("Expected body %s, got %s", tt.expected, body)
            }
        })
    }
}

// Format Check
func TestWriteError(t *testing.T) {
    tests := []struct {
        name       string
        statusCode int
        err        error
        expected   string
    }{
        {
            name:       "bad request error",
            statusCode: 400,
            err:        errors.New("invalid input"),
            expected:   `{"error":"invalid input"}`,
        },
        {
            name:       "internal server error",
            statusCode: 500,
            err:        errors.New("database connection failed"),
            expected:   `{"error":"database connection failed"}`,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            w := httptest.NewRecorder()
            
            writeError(w, tt.statusCode, tt.err)
            
            if w.Code != tt.statusCode {
                t.Errorf("Expected status %d, got %d", tt.statusCode, w.Code)
            }
            
            body := strings.TrimSpace(w.Body.String())
            if body != tt.expected {
                t.Errorf("Expected body %s, got %s", tt.expected, body)
            }
        })
    }
}