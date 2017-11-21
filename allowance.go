package allowance

import (
	"time"

	"github.com/pborman/uuid"
)

// User represents application user data model
type User struct {
	ID        uuid.UUID `storm:"id" json:"id"`
	Name      string    `json:"name"`
	Username  string    `storm:"index" json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserService interface{}

type Company struct{}
type CompanyService struct{}

type Manager struct{}
type ManagerService interface{}

type Transaction struct{}
type TransactionService interface{}
