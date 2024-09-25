package service

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-auth/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-auth/repository"
	utilsAuth "github.com/mauriciomartinezc/real-estate-mc-auth/utils"
)

type UserService interface {
	RegisterUser(user *domain.User) error
	Login(email string, password string) (*domain.User, string, error)
	Find(uuid uuid.UUID) (*domain.User, error)
	UpdateProfileId(user *domain.User, profile *domain.Profile) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (s *userService) RegisterUser(user *domain.User) error {
	userEmail, err := s.userRepository.FindByEmail(user.Email, false)
	if err != nil {
		return err
	}
	if userEmail != nil {
		return errors.New(locales.EmailAlreadyRegistered)
	}
	return s.userRepository.Create(user)
}

func (s *userService) Login(email string, password string) (*domain.User, string, error) {
	user, err := s.userRepository.FindByEmail(email, false)
	if user == nil {
		return nil, "", errors.New(locales.InvalidEmailOrPassword)
	}
	if err != nil {
		return nil, "", err
	}

	validatePassword := s.userRepository.CheckPasswordHash(password, user.Password)
	if !validatePassword {
		return &domain.User{}, "", errors.New(locales.InvalidEmailOrPassword)
	}

	token, err := utilsAuth.GenerateToken(user)

	if err != nil {
		return &domain.User{}, "", errors.New(locales.CouldNotGenerateToken)
	}

	user.Password = ""

	return user, token, nil
}

func (s *userService) Find(uuid uuid.UUID) (*domain.User, error) {
	user, err := s.userRepository.Find(uuid, "Profile")
	if user == nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) UpdateProfileId(user *domain.User, profile *domain.Profile) error {
	return s.userRepository.UpdateProfileId(user, profile)
}
