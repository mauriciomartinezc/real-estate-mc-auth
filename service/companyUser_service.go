package service

import (
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-auth/repository"
)

type CompanyUserService interface {
	AddUserToCompany(userAuth domain.User, user *domain.User, company *domain.Company, roles domain.Roles) error
	SyncUserRolesInCompany(user domain.User, company *domain.Company, roles domain.Roles) error
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

func (s companyUserService) SyncUserRolesInCompany(user domain.User, company *domain.Company, roles domain.Roles) error {
	return s.companyUserRepository.SyncUserRolesInCompany(user, company, roles)
}
