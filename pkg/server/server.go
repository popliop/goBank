package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/popliop/gobank/pkg/database"
)

type APIServer struct {
	listenAddr string
	store      database.Storage
}

func NewAPIServer(listenAddr string, store database.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	s.registerRoutes(router)

	fmt.Println("JSON API server running on port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) registerRoutes(router *mux.Router) {
	router.HandleFunc("/account", handleWrapper(s.handleAccount)).Methods("GET", "POST")
	router.HandleFunc("/account/{id}", handleWrapper(s.handleAccountByID)).Methods("GET", "DELETE")
	router.HandleFunc("/transfer", handleWrapper(s.handleTransfer)).Methods("POST")
}
