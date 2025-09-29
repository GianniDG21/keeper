// in utils_test.go
package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

// newTestRequestWithID creates a new HTTP request with chi URL params for testing purposes.
// It sets up a chi route context with the provided ID parameter and returns a request
// that can be used to test URL parameter extraction functions.
func newTestRequestWithID(id string) *http.Request {
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", id)

	req := httptest.NewRequest("GET", "/test/"+id, nil)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	return req
}

// TestGetIDFromURL tests the getIDFromURL function with various input scenarios.
// It verifies that the function correctly parses valid IDs, handles invalid inputs,
// and returns appropriate errors when expected.
func TestGetIDFromURL(t *testing.T) {
	// Test cases covering valid IDs, invalid formats, and edge cases
	testCases := []struct {
		name    string
		inputID string
		wantID  int
		wantErr bool
	}{
		{
			name:    "Valid ID",
			inputID: "123",
			wantID:  123,
			wantErr: false,
		},
		{
			name:    "Invalid ID - Not a number",
			inputID: "abc",
			wantID:  0,
			wantErr: true,
		},
		{
			name:    "Invalid ID - Empty string",
			inputID: "",
			wantID:  0,
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := newTestRequestWithID(tc.inputID)

			gotID, err := getIDFromURL(req)

			// Verify error expectation matches actual result
			if (err != nil) != tc.wantErr {
				t.Fatalf("getIDFromURL() error = %v, wantErr %v", err, tc.wantErr)
			}

			// Only check ID value when no error is expected
			if !tc.wantErr && gotID != tc.wantID {
				t.Errorf("getIDFromURL() = %v, want %v", gotID, tc.wantID)
			}
		})
	}
}