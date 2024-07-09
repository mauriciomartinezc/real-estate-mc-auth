package repository

import (
	"github.com/google/uuid"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *domain.User) error
	FindByEmail(email string) (*domain.User, error)
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
	Find(uuid uuid.UUID) (*domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (r *userRepository) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (r *userRepository) Create(user *domain.User) error {
	hashedPassword, err := r.HashPassword(user.Password)

	if err != nil {
		return err
	}

	user.Password = hashedPassword

	return r.db.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Preload("Roles.Permissions").Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Find(uuid uuid.UUID) (*domain.User, error) {
	var user domain.User
	err := r.db.Preload("Roles.Permissions").Where("id = ?", uuid).First(&user).Error
	if err != nil {
		return nil, err
	}
	user.Password = ""
	return &user, nil
}
