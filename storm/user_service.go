package storm

import (
	"allowance"

	"github.com/asdine/storm"
	"github.com/pborman/uuid"
)

// UserService represents a client to the underlying BoltDB data store.
type UserService struct {
	db *storm.DB
}

// NewUserService returns a new instance of UserService
func NewUserService() *UserService {
	return &UserService{}
}

// User returns an existing user from DB
func (s *UserService) User(id uuid.UUID) (*allowance.User, error) {
	var user *allowance.User
	err := s.db.One("ID", id.String(), &user)
	return user, err
}

// CreateUser saves a new user record to db
func (s *UserService) CreateUser(user *allowance.User) error {
	err := s.db.Save(&user)
	return err
}

// UpdateUser updates an existing user record
func (s *UserService) UpdateUser(user *allowance.User) error {
	err := s.db.Update(&user)
	return err
}

// DeleteUser removes an existing user record
func (s *UserService) DeleteUser(id uuid.UUID) error {
	u := &allowance.User{ID: id}
	err := s.db.DeleteStruct(&u)
	return err
}
