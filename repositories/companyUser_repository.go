package repositories

import (
	"errors"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-auth/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-common/cache"
	"gorm.io/gorm"
)

type CompanyUserRepository interface {
	AddUserToCompany(userAuth domain.User, user *domain.User, company *domain.Company, roles domain.Roles) error
	SyncUserRolesInCompany(userAuth domain.User, companyUser *domain.CompanyUser, roles domain.Roles) error
	FindById(id string, preloads ...string) (*domain.CompanyUser, error)
	Delete(companyUser *domain.CompanyUser) error
}

type companyUserRepository struct {
	db    *gorm.DB
	cache cache.Cache
}

func NewCompanyUserRepository(db *gorm.DB, cache cache.Cache) CompanyUserRepository {
	return &companyUserRepository{db: db, cache: cache}
}

func (r *companyUserRepository) AddUserToCompany(userAuth domain.User, user *domain.User, company *domain.Company, roles domain.Roles) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var existingCompanyUser domain.CompanyUser

		err := tx.Where("user_id = ? AND company_id = ?", user.ID, company.ID).First(&existingCompanyUser).Error
		if err == nil {
			return errors.New(locales.UserAlreadyAssociatedCompany)
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		companyUser := domain.CompanyUser{
			CreatorId: userAuth.ID.String(),
			UserId:    user.ID.String(),
			CompanyId: company.ID.String(),
			UpdaterId: nil,
		}

		if err := tx.Create(&companyUser).Error; err != nil {
			return err
		}

		if err := addRolesInBatch(tx, companyUser, roles); err != nil {
			return err
		}

		return nil
	})
}

func (r companyUserRepository) SyncUserRolesInCompany(userAuth domain.User, companyUser *domain.CompanyUser, roles domain.Roles) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		updaterId := userAuth.ID.String()
		companyUser.UpdaterId = &updaterId

		if err := tx.Save(companyUser).Error; err != nil {
			return err
		}

		var existingRoles []domain.CompanyUserRole
		if err := tx.Where("company_user_id = ?", companyUser.ID).Find(&existingRoles).Error; err != nil {
			return err
		}

		if err := syncRoles(tx, *companyUser, roles, existingRoles); err != nil {
			return err
		}

		return nil
	})
}

func (r companyUserRepository) FindById(id string, preloads ...string) (*domain.CompanyUser, error) {
	var companyUser domain.CompanyUser

	query := r.db.Where("id = ?", id)

	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	if err := query.First(&companyUser).Error; err != nil {
		return &companyUser, err
	}

	return &companyUser, nil
}

func (r companyUserRepository) Delete(companyUser *domain.CompanyUser) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var existingRoles []domain.CompanyUserRole

		if err := tx.Where("company_user_id = ?", companyUser.ID).Delete(&existingRoles).Error; err != nil {
			return err
		}

		if err := tx.Where("id = ?", companyUser.ID).Delete(&companyUser).Error; err != nil {
			return err
		}

		return nil
	})
}

func addRolesInBatch(tx *gorm.DB, companyUser domain.CompanyUser, roles domain.Roles) error {
	var existingRoles []domain.CompanyUserRole
	if err := tx.Where("company_user_id = ?", companyUser.ID.String()).Find(&existingRoles).Error; err != nil {
		return err
	}

	existingRoleMap := make(map[string]bool)
	for _, role := range existingRoles {
		existingRoleMap[role.RoleId] = true
	}

	var companyUserRoles []domain.CompanyUserRole

	for _, role := range roles {
		if !existingRoleMap[role.ID.String()] {
			companyUserRoles = append(companyUserRoles, domain.CompanyUserRole{
				CompanyUserId: companyUser.ID.String(),
				RoleId:        role.ID.String(),
			})
		}
	}

	if len(companyUserRoles) > 0 {
		if err := tx.Create(&companyUserRoles).Error; err != nil {
			return err
		}
	}

	return nil
}

func syncRoles(tx *gorm.DB, companyUser domain.CompanyUser, newRoles domain.Roles, existingRoles domain.CompanyUserRoles) error {
	existingRoleMap := make(map[string]bool)
	for _, role := range existingRoles {
		existingRoleMap[role.RoleId] = true
	}

	newRoleMap := make(map[string]bool)
	for _, role := range newRoles {
		newRoleMap[role.ID.String()] = true
		if !existingRoleMap[role.ID.String()] {
			newRole := domain.CompanyUserRole{
				CompanyUserId: companyUser.ID.String(),
				RoleId:        role.ID.String(),
			}
			if err := tx.Create(&newRole).Error; err != nil {
				return err
			}
		}
	}

	for _, existingRole := range existingRoles {
		if !newRoleMap[existingRole.RoleId] {
			if err := tx.Where("company_user_id = ? AND role_id = ?", companyUser.ID.String(), existingRole.RoleId).
				Delete(&domain.CompanyUserRole{}).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
