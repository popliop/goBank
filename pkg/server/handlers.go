package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/popliop/gobank/pkg/types"
	"github.com/popliop/gobank/pkg/utils"
)

// Wrapper Function
type apiFunc func(w http.ResponseWriter, r *http.Request) error

func handleWrapper(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			utils.WriteJSON(w, http.StatusBadRequest, utils.ApiError{Error: err.Error()})
		}
	}
}

// Handle requests to /account
func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetAccount(w, r)
	case "POST":
		return s.handleCreateAccount(w, r)
	default:
		utils.HandleNotAllowed(w, r)
		return nil
	}
}

// Handle requests to /account/{id}
func (s *APIServer) handleAccountByID(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetAccountByID(w, r)
	case "DELETE":
		return s.handleDeleteAccount(w, r)
	default:
		utils.HandleNotAllowed(w, r)
		return nil
	}
}

// GET /account
func (s *APIServer) handleGetAccount(w http.ResponseWriter, _ *http.Request) error {
	accounts, err := s.store.GetAccounts()
	if err != nil {
		return err
	}
	return utils.WriteJSON(w, http.StatusOK, accounts)
}

// POST /account
func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	createAccountReq := &types.CreateAccountRequest{}
	if err := json.NewDecoder(r.Body).Decode(createAccountReq); err != nil {
		return err
	}

	account := types.NewAccount(createAccountReq.Firstname, createAccountReq.LastName)
	if err := s.store.CreateAccount(account); err != nil {
		return err
	}
	return utils.WriteJSON(w, http.StatusOK, account)
}

// GET /account/{id}
func (s *APIServer) handleGetAccountByID(w http.ResponseWriter, r *http.Request) error {
	id, err := utils.GetID(r)
	if err != nil {
		return err
	}

	account, err := s.store.GetAccountByID(id)
	if err != nil {
		return err
	}
	return utils.WriteJSON(w, http.StatusOK, account)
}

// DELETE /account/{id}
func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("invalid id given %s", idStr)
	}

	if err := s.store.DeleteAccount(id); err != nil {
		return err
	}
	return utils.WriteJSON(w, http.StatusOK, map[string]int{"status": id})
}

// POST /transfer
func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	transferRequest := new(types.TransferRequest)
	if err := json.NewDecoder(r.Body).Decode(transferRequest); err != nil {
		return err
	}
	defer r.Body.Close()
	return utils.WriteJSON(w, http.StatusOK, transferRequest)
}
