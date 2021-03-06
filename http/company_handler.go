package http

import (
	"allowance"
	"allowance/storm"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
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

	h.Handle("/company", http.HandlerFunc(h.handleCreateCompany)).Methods("POST")
	h.Handle("/company/{id}", http.HandlerFunc(h.handleGetCompany)).Methods("GET")
	h.Handle("/company/{id}", http.HandlerFunc(h.handleUpdateCompany)).Methods("PUT")
	h.Handle("/company/{id}", http.HandlerFunc(h.handleDeleteCompany)).Methods("DELETE")

	return h
}

func (h *CompanyHandler) handleCreateCompany(w http.ResponseWriter, r *http.Request) {
	var c allowance.Company
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		EncodeError(w, err, 400, h.Logger)
	} else if err := h.CompanyService.CreateCompany(c); err != nil {
		EncodeError(w, err, 400, h.Logger)
	} else {
		EncodeJSON(w, &ResponseTemplate{Message: "success"}, h.Logger)
	}
}

func (h *CompanyHandler) handleGetCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cID := uuid.FromStringOrNil(vars["id"])

	c, err := h.CompanyService.Company(cID)
	if err != nil {
		EncodeError(w, err, 400, h.Logger)
	} else {
		EncodeJSON(w, &ResponseTemplate{Message: "success", Data: c}, h.Logger)
	}
}

func (h *CompanyHandler) handleUpdateCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cID := uuid.FromStringOrNil(vars["id"])

	var c allowance.Company
	c.ID = cID

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		EncodeError(w, err, 400, h.Logger)
	} else if err := h.CompanyService.UpdateCompany(c); err != nil {
		EncodeError(w, err, 400, h.Logger)
	} else {
		EncodeJSON(w, &ResponseTemplate{Message: "success"}, h.Logger)
	}
}

func (h *CompanyHandler) handleDeleteCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cID := uuid.FromStringOrNil(vars["id"])

	if err := h.CompanyService.DeleteCompany(cID); err != nil {
		EncodeError(w, err, 400, h.Logger)
	} else {
		EncodeJSON(w, &ResponseTemplate{Message: "success"}, h.Logger)
	}
}
