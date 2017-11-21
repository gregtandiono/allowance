package storm_test

import (
	"allowance"
	"log"
	"testing"

	"allowance/storm"

	sysstorm "github.com/asdine/storm"

	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	userService *storm.UserService
	userID_1    uuid.UUID
	userID_2    uuid.UUID
	userID_3    uuid.UUID
}

func (suite *UserServiceTestSuite) SetupSuite() {
	suite.userService = storm.NewUserService(storm.NewClient("allowance_test.db"))
	suite.userID_1 = uuid.FromStringOrNil("099ef5d7-04d2-43b0-a765-907216f388da")
	suite.userID_2 = uuid.FromStringOrNil("028b5c04-f91e-4312-990d-33525456d1a3")
	suite.userID_3 = uuid.FromStringOrNil("c509148d-2969-4b42-aa96-fa33d9884ada")

	suite.userService.Open()
	defer suite.userService.Close()

	suite.userService.CreateUser(&allowance.User{
		ID:       suite.userID_2,
		Name:     "Augustus Kwok",
		Username: "akwok",
		Password: "superdupermart",
	})

	suite.userService.CreateUser(&allowance.User{
		ID:       suite.userID_3,
		Name:     "Catherine Halim",
		Username: "chalim",
		Password: "fallout4ismyfavouritegame",
	})
}

func (suite *UserServiceTestSuite) TearDownSuite() {
	db, err := sysstorm.Open("allowance_test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.Drop("User")
}

func (suite *UserServiceTestSuite) TestUserService_CreateUser() {
	suite.userService.Open()
	defer suite.userService.Close()

	err := suite.userService.CreateUser(&allowance.User{
		ID:       suite.userID_1,
		Name:     "Gregory Tandiono",
		Username: "gtandiono",
		Password: "superawesomelongpassword",
	})
	suite.Nil(err)
}

func (suite *UserServiceTestSuite) TestUserService_CreateUser_VerifyCreate() {
	suite.userService.Open()
	defer suite.userService.Close()

	u, err := suite.userService.User(suite.userID_1)
	suite.Nil(err)
	suite.Equal("Gregory Tandiono", u.Name)
	suite.Equal("gtandiono", u.Username)
}

func (suite *UserServiceTestSuite) TestUserService_UpdateUser() {
	suite.userService.Open()
	defer suite.userService.Close()

	err := suite.userService.UpdateUser(&allowance.User{
		ID:   suite.userID_1,
		Name: "Benjamin Tandiono",
	})

	suite.Nil(err)
}

func (suite *UserServiceTestSuite) TestUserService_UpdateUser_VerifyUpdate() {
	suite.userService.Open()
	defer suite.userService.Close()

	u, err := suite.userService.User(suite.userID_1)
	suite.Nil(err)
	suite.Equal("Benjamin Tandiono", u.Name)
	suite.Equal("gtandiono", u.Username)
}

func (suite *UserServiceTestSuite) TestUserService_RemoveUser() {
	suite.userService.Open()
	defer suite.userService.Close()

	err := suite.userService.DeleteUser(suite.userID_2)
	suite.Nil(err)
}

func (suite *UserServiceTestSuite) TestUserService_RemoveUser_VerifyRemoval() {
	suite.userService.Open()
	defer suite.userService.Close()

	_, err := suite.userService.User(suite.userID_2)
	suite.NotNil(err)
	suite.Equal("not found", err.Error())
}

func TestUserServiceSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
