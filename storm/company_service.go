package storm

import (
	"allowance"
	"time"

	"github.com/satori/go.uuid"
)

// CompanyService represents a client to the underlying BoltDB data store.
type CompanyService struct {
	*Client
}

// NewCompanyService returns a new instance of CompanyService
func NewCompanyService(client *Client) *CompanyService {
	return &CompanyService{Client: client}
}

// CreateCompany saves a new company record to the DB
func (s *CompanyService) CreateCompany(c allowance.Company) error {
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()

	err := s.db.Save(&c)
	return err
}

// Company fetches an existing company record
func (s *CompanyService) Company(companyID uuid.UUID) (allowance.Company, error) {
	var company allowance.Company
	err := s.db.One("ID", companyID, &company)
	return company, err
}

// UpdateCompany updates an existing company record
func (s *CompanyService) UpdateCompany(c allowance.Company) error {
	c.UpdatedAt = time.Now()
	err := s.db.Update(&c)
	return err
}

// DeleteCompany deletes an existing company record from boltDB
func (s *CompanyService) DeleteCompany(companyID uuid.UUID) error {
	c := &allowance.Company{ID: companyID}
	err := s.db.DeleteStruct(c)
	return err
}
