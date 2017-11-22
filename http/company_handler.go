package http

import (
	"allowance/storm"
	"log"
	"os"

	"github.com/gorilla/mux"
)

// CompanyHandler represents company service REST interface
type CompanyHandler struct {
	*mux.Router
	CompanyService *storm.CompanyService
	Logger         *log.Logger
}

// NewCompanyHandler returns a new instance of UserHandler
func NewCompanyHandler(dbName string) *CompanyHandler {
	h := &CompanyHandler{
		Router:         mux.NewRouter(),
		Logger:         log.New(os.Stderr, "", log.LstdFlags),
		CompanyService: storm.NewCompanyService(storm.NewClient(dbName)),
	}

	// h.Handle("/users", http.HandlerFunc(h.handleCreateUser)).Methods("POST")
	// h.Handle("/users/{id}", http.HandlerFunc(h.handleGetUser)).Methods("GET")
	// h.Handle("/users/{id}", http.HandlerFunc(h.handleUpdateUser)).Methods("PUT")
	// h.Handle("/users/{id}", http.HandlerFunc(h.handleDeleteUser)).Methods("DELETE")

	return h
}
