package service

import (
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-auth/repository"
)

type CompanyUserService interface {
	AddUserToCompany(userAuth domain.User, user *domain.User, company *domain.Company, roles domain.Roles) error
	SyncUserRolesInCompany(userAuth domain.User, companyUser *domain.CompanyUser, roles domain.Roles) error
	FindById(companyUserId string, preloads ...string) (*domain.CompanyUser, error)
	Delete(companyUser *domain.CompanyUser) error
}

type companyUserService struct {
	companyUserRepository repository.CompanyUserRepository
}

func NewCompanyUserService(companyUserRepository repository.CompanyUserRepository) CompanyUserService {
	return &companyUserService{companyUserRepository: companyUserRepository}
}

func (s companyUserService) AddUserToCompany(userAuth domain.User, user *domain.User, company *domain.Company, roles domain.Roles) error {
	return s.companyUserRepository.AddUserToCompany(userAuth, user, company, roles)
}

func (s companyUserService) SyncUserRolesInCompany(userAuth domain.User, companyUser *domain.CompanyUser, roles domain.Roles) error {
	return s.companyUserRepository.SyncUserRolesInCompany(userAuth, companyUser, roles)
}

func (s companyUserService) FindById(companyUserId string, preloads ...string) (companyUser *domain.CompanyUser, err error) {
	return s.companyUserRepository.FindById(companyUserId, preloads...)
}

func (s companyUserService) Delete(companyUser *domain.CompanyUser) error {
	return s.companyUserRepository.Delete(companyUser)
}
