package storm

import "github.com/asdine/storm"

// UserService represents a client to the underlying BoltDB data store.
type UserService struct {
	db *storm.DB
}
