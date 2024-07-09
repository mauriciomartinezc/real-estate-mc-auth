package service

import (
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-auth/repository"
)

type RoleService interface {
	CreateRole(role *domain.Role) error
	// Otros métodos necesarios
}

type roleService struct {
	roleRepository repository.RoleRepository
}

func NewRoleService(roleRepo repository.RoleRepository) RoleService {
	return &roleService{
		roleRepository: roleRepo,
	}
}

func (s *roleService) CreateRole(role *domain.Role) error {
	return s.roleRepository.Create(role)
}
