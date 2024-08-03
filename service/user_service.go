package service

import (
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-auth/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-auth/repository"
	utilsAuth "github.com/mauriciomartinezc/real-estate-mc-auth/utils"
)

type UserService interface {
	RegisterUser(user *domain.User) error
	Login(email string, password string) (*domain.User, string, error)
	GetMe(e echo.Context) (*domain.User, error)
}

type userService struct {
	userRepository       repository.UserRepository
	roleRepository       repository.RoleRepository
	permissionRepository repository.PermissionRepository
}

func NewUserService(userRepo repository.UserRepository, roleRepo repository.RoleRepository, permRepo repository.PermissionRepository) UserService {
	return &userService{
		userRepository:       userRepo,
		roleRepository:       roleRepo,
		permissionRepository: permRepo,
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

func (s *userService) GetMe(c echo.Context) (*domain.User, error) {
	userId, ok := c.Get("userId").(uuid.UUID)
	if !ok {
		return &domain.User{}, errors.New(locales.CouldNotGetUserId)
	}

	return s.userRepository.Find(userId, "Roles.Permissions")
}
