package allowance

import (
	"time"

	"github.com/satori/go.uuid"
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

// UserService represents User model CRUD interface against the boltDB
type UserService interface {
	User(userID uuid.UUID) (*User, error)
	CreateUser(u *User) error
	UpdateUser(u *User) error
	DeleteUser(userID uuid.UUID) error
}

// Company represents a company registered in this application
type Company struct {
	ID        uuid.UUID `storm:"id" json:"id"`
	Name      string    `storm:"index" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CompanyService represents company model CRUD interface against the boltDB
type CompanyService interface {
	Company(companyID uuid.UUID) (Company, error)
	CreateCompany(c Company) error
	UpdateCompany(c Company) error
	DeleteCompany(companyID uuid.UUID) error
}

type Manager struct{}
type ManagerService interface{}

type Transaction struct{}
type TransactionService interface{}
