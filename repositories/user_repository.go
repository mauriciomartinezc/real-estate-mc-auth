package repositories

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-auth/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-common/cache"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *domain.User) error
	FindByEmail(email string, opts ...bool) (*domain.User, error)
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
	Find(uuid uuid.UUID, preload ...string) (*domain.User, error)
	UpdateProfileId(user *domain.User, profile *domain.Profile) error
	ResetPassword(user domain.User, oldPassword string, newPassword string) error
}

type userRepository struct {
	db    *gorm.DB
	cache cache.Cache
}

func NewUserRepository(db *gorm.DB, cache cache.Cache) UserRepository {
	return &userRepository{db: db, cache: cache}
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

	if user.ProfileId == nil {
		profileRepository := NewProfileRepository(r.db, r.cache)
		profile, err := profileRepository.Create(user, new(domain.Profile))
		if err != nil {
			return err
		}
		profileId := profile.ID.String()
		user.ProfileId = &profileId
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

func (r *userRepository) Find(uuid uuid.UUID, preloads ...string) (*domain.User, error) {
	var user domain.User

	query := r.db
	for _, preload := range preloads {
		if preload != "" {
			query = query.Preload(preload)
		}
	}

	err := query.Where("id = ?", uuid).First(&user).Error
	if err != nil {
		return nil, err
	}
	user.Password = ""
	return &user, nil
}

func (r *userRepository) UpdateProfileId(user *domain.User, profile *domain.Profile) error {
	err := r.db.Model(user).Update("ProfileId", profile.ID).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) ResetPassword(user domain.User, oldPassword string, newPassword string) error {
	r.db.Where("id = ?", user.ID).First(&user)

	validatePassword := r.CheckPasswordHash(oldPassword, user.Password)

	if !validatePassword {
		return errors.New(locales.InvalidOldPassword)
	}

	hashedPassword, err := r.HashPassword(newPassword)

	if err != nil {
		return err
	}

	user.Password = hashedPassword
	user.RequiredResetPassword = false
	return r.db.Save(&user).Error
}
