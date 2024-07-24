package repository

import (
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(role *domain.Role) error
	FindBySlug(slug string) (*domain.Role, error)
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

func (r *roleRepository) FindBySlug(slug string) (*domain.Role, error) {
	var role domain.Role
	err := r.db.Where("slug = ?", slug).First(&role).Error
	return &role, err
}
