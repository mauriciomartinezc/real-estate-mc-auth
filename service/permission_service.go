package service

import (
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-auth/repository"
)

type PermissionService interface {
	CreatePermission(permission *domain.Permission) error
	// Otros métodos necesarios
}

type permissionService struct {
	permissionRepository repository.PermissionRepository
}

func NewPermissionService(permissionRepo repository.PermissionRepository) PermissionService {
	return &permissionService{
		permissionRepository: permissionRepo,
	}
}

func (s *permissionService) CreatePermission(permission *domain.Permission) error {
	return s.permissionRepository.Create(permission)
}
