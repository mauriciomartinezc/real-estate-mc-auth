package repositories

import (
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-common/cache"
	"gorm.io/gorm"
)

type PermissionRepository interface {
	Create(permission *domain.Permission) error
	// Otros métodos necesarios
}

type permissionRepository struct {
	db    *gorm.DB
	cache cache.Cache
}

func NewPermissionRepository(db *gorm.DB, cache cache.Cache) PermissionRepository {
	return &permissionRepository{db: db, cache: cache}
}

func (r *permissionRepository) Create(permission *domain.Permission) error {
	return r.db.Create(permission).Error
}
