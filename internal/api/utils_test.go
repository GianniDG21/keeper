//in utils_test.go
package api

import (
	"context"
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/go-chi/chi/v5"
)

func newTestRequestWithID(id string) *http.Request {	// Helper function to create a new HTTP request with chi URL params
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", id)

	req := httptest.NewRequest("GET", "/test/"+id, nil)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	return req
}

func TestGetIDFromURL(t *testing.T) {
	//Test cases
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

			if (err != nil) != tc.wantErr {
				t.Fatalf("getIDFromURL() error = %v, wantErr %v", err, tc.wantErr)
			}
			
			if !tc.wantErr && gotID != tc.wantID {
				t.Errorf("getIDFromURL() = %v, want %v", gotID, tc.wantID)
			}
		})
	}
}