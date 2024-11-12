package repositories

import (
	"github.com/google/uuid"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-common/cache"
	"gorm.io/gorm"
)

type ProfileRepository interface {
	Create(user *domain.User, profile *domain.Profile) (*domain.Profile, error)
	Update(profile *domain.Profile) (*domain.Profile, error)
	FindByUserId(userId uuid.UUID) (*domain.Profile, error)
}

type profileRepository struct {
	db    *gorm.DB
	cache cache.Cache
}

func NewProfileRepository(db *gorm.DB, cache cache.Cache) ProfileRepository {
	return &profileRepository{db: db, cache: cache}
}

func (p profileRepository) Create(_ *domain.User, profile *domain.Profile) (*domain.Profile, error) {
	if err := p.db.Preload("City").Create(profile).Error; err != nil {
		return nil, err
	}
	return profile, nil
}

func (p profileRepository) Update(profile *domain.Profile) (*domain.Profile, error) {
	if err := p.db.Preload("City").Save(profile).Error; err != nil {
		return nil, err
	}
	return profile, nil
}

func (p profileRepository) FindByUserId(userId uuid.UUID) (*domain.Profile, error) {
	userRepo := NewUserRepository(p.db, p.cache)
	user, err := userRepo.Find(userId, "Profile")
	if err != nil {
		return nil, err
	}

	return user.Profile, nil
}
