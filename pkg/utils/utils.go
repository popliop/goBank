package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ApiError represents a structured error response
type ApiError struct {
	Error string `json:"error"`
}

// WriteJSON writes a JSON response to the client
func WriteJSON(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// GetID extracts and validates the ID parameter from the request
func GetID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("invalid ID given: %s", idStr)
	}
	return id, nil
}

// HandleNotAllowed writes a "method not allowed" response
func HandleNotAllowed(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Unsupported method %s called on %s\n", r.Method, r.URL.Path)
	WriteJSON(w, http.StatusMethodNotAllowed, ApiError{Error: "Method not allowed"})
}
