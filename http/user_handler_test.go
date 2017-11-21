package http_test

import (
	"allowance"
	ahttp "allowance/http"
	"allowance/storm"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	sysstorm "github.com/asdine/storm"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/suite"
)

type getUserResponse struct {
	Message string
	Error   string
	Data    *allowance.User
}

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

func (suite *UserHandlerTestSuite) TestUserHandler_FetchUser() {
	request, _ := http.NewRequest("GET", "/users/"+suite.userID_2.String(), nil)
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	h := suite.userHandler
	h.UserService.Open()
	defer h.UserService.Close()

	h.ServeHTTP(response, request)

	var responseBody *getUserResponse
	json.Unmarshal(response.Body.Bytes(), &responseBody)

	suite.Equal("", responseBody.Error)
	suite.Equal("success", responseBody.Message)
	suite.Equal("Augustus Kwok", responseBody.Data.Name)
	suite.Equal("akwok", responseBody.Data.Username)
}

func (suite *UserHandlerTestSuite) TestUserHandler_CreateUser() {
	mockData := []byte(`{
		"id": "` + suite.userID_1.String() + `",
		"name": "Gregory Tandiono",
		"username": "gtandiono",
		"password": "somesuperawesomepassword"
	}`)
	request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(mockData))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	h := suite.userHandler
	h.UserService.Open()
	defer h.UserService.Close()

	h.ServeHTTP(response, request)

	var responseBody *ahttp.ResponseTemplate
	json.Unmarshal(response.Body.Bytes(), &responseBody)

	suite.Equal("", responseBody.Error)
	suite.Equal("success", responseBody.Message)
}

func (suite *UserHandlerTestSuite) TestUserHandler_UpdateUser() {
	mockData := []byte(`{
		"name": "Jupiter Grog",
		"username": "jgrog"
	}`)
	request, _ := http.NewRequest("PUT", "/users/"+suite.userID_2.String(), bytes.NewBuffer(mockData))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	h := suite.userHandler
	h.UserService.Open()
	defer h.UserService.Close()

	h.ServeHTTP(response, request)

	var responseBody *ahttp.ResponseTemplate
	json.Unmarshal(response.Body.Bytes(), &responseBody)

	suite.Equal("", responseBody.Error)
	suite.Equal("success", responseBody.Message)
}

func (suite *UserHandlerTestSuite) TestUserHandler_UpdateUser_VerifyUpdate() {
	request, _ := http.NewRequest("GET", "/users/"+suite.userID_2.String(), nil)
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	h := suite.userHandler
	h.UserService.Open()
	defer h.UserService.Close()

	h.ServeHTTP(response, request)

	var responseBody *getUserResponse
	json.Unmarshal(response.Body.Bytes(), &responseBody)

	suite.Equal("", responseBody.Error)
	suite.Equal("success", responseBody.Message)
	suite.Equal("Jupiter Grog", responseBody.Data.Name)
	suite.Equal("jgrog", responseBody.Data.Username)
}

func (suite *UserHandlerTestSuite) TestUserHandler_DeleteUser() {
	request, _ := http.NewRequest("DELETE", "/users/"+suite.userID_3.String(), nil)
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	h := suite.userHandler
	h.UserService.Open()
	defer h.UserService.Close()

	h.ServeHTTP(response, request)

	var responseBody *ahttp.ResponseTemplate
	json.Unmarshal(response.Body.Bytes(), &responseBody)

	suite.Equal("", responseBody.Error)
	suite.Equal("success", responseBody.Message)
}

func (suite *UserHandlerTestSuite) TestUserHandler_DeleteUser_VerifyDelete() {
	request, _ := http.NewRequest("GET", "/users/"+suite.userID_3.String(), nil)
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	h := suite.userHandler
	h.UserService.Open()
	defer h.UserService.Close()

	h.ServeHTTP(response, request)

	var responseBody *getUserResponse
	json.Unmarshal(response.Body.Bytes(), &responseBody)

	suite.Equal("not found", responseBody.Error)
	suite.Equal("fail", responseBody.Message)
}

func TestUserHandlerSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}
