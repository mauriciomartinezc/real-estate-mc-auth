package services

import (
	"github.com/google/uuid"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-auth/repositories"
)

type CompanyService interface {
	Create(company *domain.Company, user domain.User) error
	FindById(id string) (*domain.Company, error)
	Update(company *domain.Company) error
	AssociateUserToCompany(company *domain.Company, user domain.User, role domain.Role, userCreator domain.User) error
	CompaniesMe(user domain.User) (domain.Companies, error)
}

type companyService struct {
	companyRepository repositories.CompanyRepository
}

func NewCompanyService(companyRepository repositories.CompanyRepository) CompanyService {
	return &companyService{companyRepository: companyRepository}
}

func (s *companyService) Create(company *domain.Company, user domain.User) error {
	return s.companyRepository.Create(company, user)
}

func (s *companyService) FindById(id string) (*domain.Company, error) {
	return s.companyRepository.FindByID(uuid.MustParse(id))
}

func (s *companyService) Update(company *domain.Company) error {
	return s.companyRepository.Update(company)
}

func (s *companyService) AssociateUserToCompany(company *domain.Company, user domain.User, role domain.Role, userCreator domain.User) error {
	return s.companyRepository.AssociateUserToCompany(company, user, role, userCreator)
}

func (s *companyService) CompaniesMe(user domain.User) (domain.Companies, error) {
	return s.companyRepository.CompaniesMe(user)
}
