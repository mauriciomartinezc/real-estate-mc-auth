package repository

import (
	"github.com/google/uuid"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"gorm.io/gorm"
)

type ProfileRepository interface {
	Create(user *domain.User, profile *domain.Profile) (*domain.Profile, error)
	Update(profile *domain.Profile) (*domain.Profile, error)
	FindByUserId(userId uuid.UUID) (*domain.Profile, error)
}

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{db: db}
}

func (p profileRepository) Create(user *domain.User, profile *domain.Profile) (*domain.Profile, error) {
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
	userRepo := NewUserRepository(p.db)
	user, err := userRepo.Find(userId, "Profile")
	if err != nil {
		return nil, err
	}

	return user.Profile, nil
}
