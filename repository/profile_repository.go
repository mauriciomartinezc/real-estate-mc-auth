package repository

import (
	"github.com/google/uuid"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"gorm.io/gorm"
)

type ProfileRepository interface {
	Create(profile *domain.Profile) (*domain.Profile, error)
	Update(profile *domain.Profile) (*domain.Profile, error)
	FindByUserId(userId uuid.UUID) (*domain.Profile, error)
}

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{db: db}
}

func (p profileRepository) Create(profile *domain.Profile) (*domain.Profile, error) {
	if err := p.db.Create(profile).Error; err != nil {
		return nil, err
	}
	return profile, nil
}

func (p profileRepository) Update(profile *domain.Profile) (*domain.Profile, error) {
	if err := p.db.Save(profile).Error; err != nil {
		return nil, err
	}
	return profile, nil
}

func (p profileRepository) FindByUserId(userId uuid.UUID) (*domain.Profile, error) {
	var profile domain.Profile
	userRepo := NewUserRepository(p.db)
	user, err := userRepo.Find(userId, "")
	if err != nil {
		return nil, err
	}
	err = p.db.Where("id = ?", user.ProfileId).First(profile).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}
