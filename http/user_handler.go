package http

import (
	"allowance/storm"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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
	return h
}

func (h *UserHandler) handleCreateUser(w http.ResponseWriter, r *http.Request) {
}

func (h *UserHandler) handleGetUser(w http.ResponseWriter, r *http.Request) {
}

func (h *UserHandler) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
}

func (h *UserHandler) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
}
