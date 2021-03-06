package storm

import (
	"github.com/asdine/storm"
)

// Client represents a storm client
type Client struct {
	db     *storm.DB
	dbName string
}

// NewClient returns a new storm client
func NewClient(dbName string) *Client {
	return &Client{dbName: dbName}
}

// Open opens a new boltdb conn
func (c *Client) Open() error {
	db, err := storm.Open(c.dbName, storm.BoltOptions(0600, nil))
	if err != nil {
		return err
	}
	c.db = db
	return nil
}

// Close closes an existing boltdb conn
func (c *Client) Close() error {
	if c.db != nil {
		return c.db.Close()
	}
	return nil
}
