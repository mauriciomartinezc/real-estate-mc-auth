package services

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-auth/repositories"
)

type ProfileService interface {
	Create(user *domain.User, profile *domain.Profile) (*domain.Profile, error)
	Update(profile *domain.Profile) (*domain.Profile, error)
	MeProfile(c echo.Context) (*domain.Profile, error)
}

type profileService struct {
	profileRepository repositories.ProfileRepository
}

func NewProfileService(profileRepository repositories.ProfileRepository) ProfileService {
	return &profileService{
		profileRepository: profileRepository,
	}
}

func (p *profileService) Create(user *domain.User, profile *domain.Profile) (*domain.Profile, error) {
	return p.profileRepository.Create(user, profile)
}

func (p *profileService) Update(profile *domain.Profile) (*domain.Profile, error) {
	return p.profileRepository.Update(profile)
}

func (p *profileService) MeProfile(c echo.Context) (*domain.Profile, error) {
	userId, _ := c.Get("userId").(uuid.UUID)
	return p.profileRepository.FindByUserId(userId)
}
