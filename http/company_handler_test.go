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

type getCompanyResponse struct {
	Message string
	Error   string
	Data    allowance.Company
}

type CompanyHandlerTestSuite struct {
	suite.Suite
	companyService *storm.CompanyService
	companyHandler *ahttp.CompanyHandler
	companyID_1    uuid.UUID
	companyID_2    uuid.UUID
	companyID_3    uuid.UUID
}

func (suite *CompanyHandlerTestSuite) SetupSuite() {
	suite.companyService = storm.NewCompanyService(storm.NewClient("allowance_test.db"))
	suite.companyHandler = ahttp.NewCompanyHandler("allowance_test.db")
	suite.companyID_1 = uuid.NewV4()
	suite.companyID_2 = uuid.NewV4()
	suite.companyID_3 = uuid.NewV4()

	suite.companyService.Open()
	defer suite.companyService.Close()

	suite.companyService.CreateCompany(allowance.Company{
		ID:   suite.companyID_2,
		Name: "Cargill Feed",
	})

	suite.companyService.CreateCompany(allowance.Company{
		ID:   suite.companyID_3,
		Name: "Chel Jedang",
	})
}

func (suite *CompanyHandlerTestSuite) TearDownSuite() {
	db, err := sysstorm.Open("allowance_test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.Drop("Company")
}

func (suite *CompanyHandlerTestSuite) TestCompanyHandler_FetchCompany() {
	request, _ := http.NewRequest("GET", "/company/"+suite.companyID_2.String(), nil)
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	h := suite.companyHandler
	h.CompanyService.Open()
	defer h.CompanyService.Close()

	h.ServeHTTP(response, request)

	var responseBody *getCompanyResponse
	json.Unmarshal(response.Body.Bytes(), &responseBody)
	suite.Equal("", responseBody.Error)
	suite.Equal("success", responseBody.Message)
	suite.Equal("Cargill Feed", responseBody.Data.Name)
}

func (suite *CompanyHandlerTestSuite) TestCompanyHandler_CreateCompany() {
	mockData := []byte(`{
		"id": "` + suite.companyID_1.String() + `",
		"name": "PT Agrix Indonesia"
	}`)
	request, _ := http.NewRequest("POST", "/company", bytes.NewBuffer(mockData))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	h := suite.companyHandler
	h.CompanyService.Open()
	defer h.CompanyService.Close()

	h.ServeHTTP(response, request)

	var responseBody *ahttp.ResponseTemplate
	json.Unmarshal(response.Body.Bytes(), &responseBody)

	suite.Empty(responseBody.Error)
	suite.Equal("success", responseBody.Message)
}

func (suite *CompanyHandlerTestSuite) TestCompanyHandler_UpdateCompany() {
	mockData := []byte(`{
		"name": "PT Agrofood Makmur"
	}`)
	request, _ := http.NewRequest("PUT", "/company/"+suite.companyID_1.String(), bytes.NewBuffer(mockData))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	h := suite.companyHandler
	h.CompanyService.Open()
	defer h.CompanyService.Close()

	h.ServeHTTP(response, request)

	var responseBody *ahttp.ResponseTemplate
	json.Unmarshal(response.Body.Bytes(), &responseBody)

	suite.Empty(responseBody.Error)
	suite.Equal("success", responseBody.Message)
}

func (suite *CompanyHandlerTestSuite) TestCompanyHandler_UpdateCompany_VerifyUpdate() {
	request, _ := http.NewRequest("GET", "/company/"+suite.companyID_1.String(), nil)
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	h := suite.companyHandler
	h.CompanyService.Open()
	defer h.CompanyService.Close()

	h.ServeHTTP(response, request)

	var responseBody *getCompanyResponse
	json.Unmarshal(response.Body.Bytes(), &responseBody)

	suite.Empty(responseBody.Error)
	suite.Equal("success", responseBody.Message)
	suite.Equal("PT Agrofood Makmur", responseBody.Data.Name)
}

func (suite *CompanyHandlerTestSuite) TestCompanyHandler_DeleteCompany() {
	request, _ := http.NewRequest("DELETE", "/company/"+suite.companyID_3.String(), nil)
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	h := suite.companyHandler
	h.CompanyService.Open()
	defer h.CompanyService.Close()

	h.ServeHTTP(response, request)

	var responseBody *getCompanyResponse
	json.Unmarshal(response.Body.Bytes(), &responseBody)

	suite.Empty(responseBody.Error)
	suite.Equal("success", responseBody.Message)
}

func (suite *CompanyHandlerTestSuite) TestCompanyHandler_DeleteCompany_VerifyDelete() {
	request, _ := http.NewRequest("GET", "/company/"+suite.companyID_3.String(), nil)
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	h := suite.companyHandler
	h.CompanyService.Open()
	defer h.CompanyService.Close()

	h.ServeHTTP(response, request)

	var responseBody *getCompanyResponse
	json.Unmarshal(response.Body.Bytes(), &responseBody)

	suite.Equal("fail", responseBody.Message)
	suite.Equal("not found", responseBody.Error)
}

func TestCompanyHandlerSuite(t *testing.T) {
	suite.Run(t, new(CompanyHandlerTestSuite))
}
