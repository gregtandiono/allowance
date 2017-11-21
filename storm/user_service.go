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

// User returns an existing user from DB
func (s *UserService) User(id uuid.UUID) *allowance.User {}
