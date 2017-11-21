package http_test

import (
	"allowance"
	ahttp "allowance/http"
	"allowance/storm"
	"log"
	"testing"

	sysstorm "github.com/asdine/storm"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/suite"
)

type UserHandlerTestSuite struct {
	suite.Suite
	userService *storm.UserService
	userHandler *ahttp.UserHandler
	userID_1    uuid.UUID
	userID_2    uuid.UUID
	userID_3    uuid.UUID
}

func (suite *UserHandlerTestSuite) SetupSuite() {
	suite.userService = storm.NewUserService(storm.NewClient("allowance_test.db"))
	suite.userHandler = ahttp.NewUserHandler("allowance_test.db")
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

func (suite *UserHandlerTestSuite) TearDownSuite() {
	db, err := sysstorm.Open("allowance_test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.Drop("User")
}

func TestUserHandlerSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}
