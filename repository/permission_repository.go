package repository

import (
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"gorm.io/gorm"
)

type PermissionRepository interface {
	Create(permission *domain.Permission) error
	// Otros métodos necesarios
}

type permissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return &permissionRepository{db: db}
}

func (r *permissionRepository) Create(permission *domain.Permission) error {
	return r.db.Create(permission).Error
}
