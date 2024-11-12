package repositories

import (
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-common/cache"
	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(role *domain.Role) error
	FindBySlug(slug string) (*domain.Role, error)
}

type roleRepository struct {
	db    *gorm.DB
	cache cache.Cache
}

func NewRoleRepository(db *gorm.DB, cache cache.Cache) RoleRepository {
	return &roleRepository{db: db, cache: cache}
}

func (r *roleRepository) Create(role *domain.Role) error {
	return r.db.Create(role).Error
}

func (r *roleRepository) FindBySlug(slug string) (*domain.Role, error) {
	var role domain.Role
	err := r.db.Where("slug = ?", slug).First(&role).Error
	return &role, err
}
