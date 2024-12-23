package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

// Helper Functions
type ApiError struct {
	Error string
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func getID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("invalid id given %s", idStr)
	}
	return id, nil
}

func handleNotAllowed(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Unsupported method %s called on %s\n", r.Method, r.URL.Path)
	WriteJSON(w, http.StatusMethodNotAllowed, ApiError{Error: "Method not allowed"})
}
