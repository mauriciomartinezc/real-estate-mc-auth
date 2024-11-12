package services

import (
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-auth/repositories"
)

type RoleService interface {
	CreateRole(role *domain.Role) error
	// Otros métodos necesarios
}

type roleService struct {
	roleRepository repositories.RoleRepository
}

func NewRoleService(roleRepo repositories.RoleRepository) RoleService {
	return &roleService{
		roleRepository: roleRepo,
	}
}

func (s *roleService) CreateRole(role *domain.Role) error {
	return s.roleRepository.Create(role)
}
