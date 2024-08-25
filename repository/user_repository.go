package repository

import (
	"github.com/google/uuid"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-common/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *domain.User) error
	FindByEmail(email string, opts ...bool) (*domain.User, error)
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
	Find(uuid uuid.UUID, preload string) (*domain.User, error)
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

	if len(user.Roles) == 0 {
		dataRole := domain.ROLES["USER"]
		role := new(domain.Role)
		role.ID = dataRole.ID
		role.Name = dataRole.Name
		role.Slug = dataRole.Slug
		user.Roles = []domain.Role{*role}
	}

	if !utils.IsValidUUID(user.ProfileId.String()) {
		profileRepository := NewProfileRepository(r.db)
		profile, err := profileRepository.Create(new(domain.Profile))
		if err != nil {
			return err
		}
		user.ProfileId = profile.ID
	}

	return r.db.Create(user).Error
}

func (r *userRepository) FindByEmail(email string, opts ...bool) (*domain.User, error) {
	requiredError := false
	if len(opts) > 0 {
		requiredError = opts[0]
	}

	var user domain.User
	query := r.db.Where("email = ?", email)
	err := query.First(&user).Error
	if requiredError && err != nil {
		return nil, err
	}
	if err != nil {
		return nil, nil
	}
	return &user, nil
}

func (r *userRepository) Find(uuid uuid.UUID, preload string) (*domain.User, error) {
	var user domain.User
	err := r.db.Preload(preload).Where("id = ?", uuid).First(&user).Error
	if err != nil {
		return nil, err
	}
	user.Password = ""
	return &user, nil
}
