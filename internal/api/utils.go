//internal/api/utils.go

package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"errors"
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