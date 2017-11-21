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

// UserHandler represents user service REST interface
type UserHandler struct {
	*mux.Router
	UserService *storm.UserService
	Logger      *log.Logger
}

// NewUserHandler returns a new instance of UserHandler
func NewUserHandler(dbName string) *UserHandler {
	h := &UserHandler{
		Router:      mux.NewRouter(),
		Logger:      log.New(os.Stderr, "", log.LstdFlags),
		UserService: storm.NewUserService(storm.NewClient(dbName)),
	}

	h.Handle("/users", http.HandlerFunc(h.handleCreateUser)).Methods("POST")
	h.Handle("/users/{id}", http.HandlerFunc(h.handleGetUser)).Methods("GET")
	h.Handle("/users/{id}", http.HandlerFunc(h.handleUpdateUser)).Methods("PUT")
	h.Handle("/users/{id}", http.HandlerFunc(h.handleDeleteUser)).Methods("DELETE")

	return h
}

func (h *UserHandler) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var user allowance.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		EncodeError(w, err, 400, h.Logger)
	} else if err := h.UserService.CreateUser(&user); err != nil {
		EncodeError(w, err, 400, h.Logger)
	} else {
		EncodeJSON(w, &ResponseTemplate{Message: "success"}, h.Logger)
	}
}

func (h *UserHandler) handleGetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := uuid.FromStringOrNil(vars["id"])

	u, err := h.UserService.User(userID)
	if err != nil {
		EncodeError(w, err, 400, h.Logger)
	} else {
		EncodeJSON(w, &ResponseTemplate{Message: "success", Data: u}, h.Logger)
	}
}

func (h *UserHandler) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := uuid.FromStringOrNil(vars["id"])

	var user allowance.User
	user.ID = userID

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		EncodeError(w, err, 400, h.Logger)
	} else if err := h.UserService.UpdateUser(&user); err != nil {
		EncodeError(w, err, 400, h.Logger)
	} else {
		EncodeJSON(w, &ResponseTemplate{Message: "success"}, h.Logger)
	}
}

func (h *UserHandler) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := uuid.FromStringOrNil(vars["id"])

	if err := h.UserService.DeleteUser(userID); err != nil {
		EncodeError(w, err, 400, h.Logger)
	} else {
		EncodeJSON(w, &ResponseTemplate{Message: "success"}, h.Logger)
	}
}
