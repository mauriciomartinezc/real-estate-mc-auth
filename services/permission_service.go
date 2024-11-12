package services

import (
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-auth/repositories"
)

type PermissionService interface {
	CreatePermission(permission *domain.Permission) error
	// Otros métodos necesarios
}

type permissionService struct {
	permissionRepository repositories.PermissionRepository
}

func NewPermissionService(permissionRepo repositories.PermissionRepository) PermissionService {
	return &permissionService{
		permissionRepository: permissionRepo,
	}
}

func (s *permissionService) CreatePermission(permission *domain.Permission) error {
	return s.permissionRepository.Create(permission)
}
