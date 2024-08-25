package service

import (
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-auth/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-auth/repository"
)

type ProfileService interface {
	Create(profile *domain.Profile) (*domain.Profile, error)
	Update(profile *domain.Profile) (*domain.Profile, error)
	MeProfile(c echo.Context) (*domain.Profile, error)
}

type profileService struct {
	profileRepository repository.ProfileRepository
}

func NewProfileService(profileRepository repository.ProfileRepository) ProfileService {
	return &profileService{
		profileRepository: profileRepository,
	}
}

func (p *profileService) Create(profile *domain.Profile) (*domain.Profile, error) {
	return p.profileRepository.Create(profile)
}

func (p *profileService) Update(profile *domain.Profile) (*domain.Profile, error) {
	return p.profileRepository.Update(profile)
}

func (p *profileService) MeProfile(c echo.Context) (*domain.Profile, error) {
	userId, ok := c.Get("userId").(uuid.UUID)
	if !ok {
		return nil, errors.New(locales.CouldNotGetUserId)
	}
	return p.profileRepository.FindByUserId(userId)
}
