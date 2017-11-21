package storm_test

import (
	"datwire/pkg/bolt"
	"testing"

	"github.com/pborman/uuid"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	userService *bolt.UserService
	userID_1    uuid.UUID
	userID_2    uuid.UUID
	userID_3    uuid.UUID
}

func (suite *UserServiceTestSuite) SetupSuite() {

}

func (suite *UserServiceTestSuite) TearDownSuite() {

}

func TestUserServiceSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
