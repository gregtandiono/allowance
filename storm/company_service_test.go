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

type CompanyServiceTestSuite struct {
	suite.Suite
	companyService *storm.CompanyService
	companyID_1    uuid.UUID
	companyID_2    uuid.UUID
	companyID_3    uuid.UUID
}

func (suite *CompanyServiceTestSuite) SetupSuite() {
	suite.companyService = storm.NewCompanyService(storm.NewClient("allowance_test.db"))
	suite.companyID_1, _ = uuid.NewV4()
	suite.companyID_2, _ = uuid.NewV4()
	suite.companyID_3, _ = uuid.NewV4()

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

func (suite *CompanyServiceTestSuite) TearDownSuite() {
	db, err := sysstorm.Open("allowance_test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.Drop("Company")
}

func (suite *CompanyServiceTestSuite) TestCompanyService_CreateCompany() {
	suite.companyService.Open()
	defer suite.companyService.Close()

	err := suite.companyService.CreateCompany(allowance.Company{
		ID:   suite.companyID_1,
		Name: "PT Agrix Indonesia",
	})

	suite.Nil(err)
}

func (suite *CompanyServiceTestSuite) TestCompanyService_CreateCompany_VerifyCreate() {
	suite.companyService.Open()
	defer suite.companyService.Close()

	c, err := suite.companyService.Company(suite.companyID_1)
	suite.Nil(err)
	suite.Equal("PT Agrix Indonesia", c.Name)
}

func (suite *CompanyServiceTestSuite) TestCompanyService_UpdateCompany() {
	suite.companyService.Open()
	defer suite.companyService.Close()

	err := suite.companyService.UpdateCompany(allowance.Company{
		ID:   suite.companyID_1,
		Name: "PT Japfa Comfeed",
	})

	suite.Nil(err)
}

func (suite *CompanyServiceTestSuite) TestCompanyService_UpdateCompany_VerifyUpdate() {
	suite.companyService.Open()
	defer suite.companyService.Close()

	c, err := suite.companyService.Company(suite.companyID_1)
	suite.Nil(err)
	suite.Equal("PT Japfa Comfeed", c.Name)
}

func (suite *CompanyServiceTestSuite) TestCompanyService_DeleteCompany() {
	suite.companyService.Open()
	defer suite.companyService.Close()

	err := suite.companyService.DeleteCompany(suite.companyID_2)
	suite.Nil(err)
}

func (suite *CompanyServiceTestSuite) TestCompanyService_DeleteCompany_VerifyDelete() {
	suite.companyService.Open()
	defer suite.companyService.Close()

	_, err := suite.companyService.Company(suite.companyID_2)
	suite.NotNil(err)
	suite.Equal("not found", err.Error())
}

func TestCustomerServiceSuite(t *testing.T) {
	suite.Run(t, new(CompanyServiceTestSuite))
}
