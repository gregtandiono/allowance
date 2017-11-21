package storm

import (
	"allowance"

	"golang.org/x/crypto/bcrypt"

	"github.com/satori/go.uuid"
)

// UserService represents a client to the underlying BoltDB data store.
type UserService struct {
	*Client
}

// NewUserService returns a new instance of UserService
func NewUserService(client *Client) *UserService {
	return &UserService{Client: client}
}

// User returns an existing user from DB
func (s *UserService) User(id uuid.UUID) (allowance.User, error) {
	var user allowance.User
	err := s.db.One("ID", id, &user)
	return user, err
}

// CreateUser saves a new user record to db
func (s *UserService) CreateUser(user *allowance.User) error {
	hp, err := s.hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hp)
	err = s.db.Save(user)
	return err
}

// UpdateUser updates an existing user record
func (s *UserService) UpdateUser(user *allowance.User) error {
	err := s.db.Update(&user)
	return err
}

// DeleteUser removes an existing user record
func (s *UserService) DeleteUser(userID uuid.UUID) error {
	u := &allowance.User{ID: userID}
	err := s.db.DeleteStruct(&u)
	return err
}

func (s *UserService) hashPassword(password string) ([]byte, error) {
	p := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(p, bcrypt.DefaultCost)
	return hash, err
}
