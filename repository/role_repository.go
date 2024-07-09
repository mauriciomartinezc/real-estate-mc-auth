package repository

import (
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(role *domain.Role) error
	// Otros métodos necesarios
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) Create(role *domain.Role) error {
	return r.db.Create(role).Error
}
