package service

import (
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-auth/repository"
	"github.com/mauriciomartinezc/real-estate-mc-auth/utils"
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
	return s.userRepository.Create(user)
}

func (s *userService) Login(email string, password string) (*domain.User, string, error) {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return nil, "", err
	}

	validatePassword := s.userRepository.CheckPasswordHash(password, user.Password)
	if !validatePassword {
		return &domain.User{}, "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(user)

	if err != nil {
		return &domain.User{}, "", errors.New("could not generate token")
	}

	user.Password = ""

	return user, token, nil
}

func (s *userService) GetMe(c echo.Context) (*domain.User, error) {
	userId, ok := c.Get("userId").(uuid.UUID)
	if !ok {
		return &domain.User{}, errors.New("could not get user id")
	}

	return s.userRepository.Find(userId)
}
