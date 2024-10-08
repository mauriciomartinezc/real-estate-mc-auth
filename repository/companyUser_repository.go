package repository

import (
	"errors"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-auth/i18n/locales"
	"gorm.io/gorm"
)

type CompanyUserRepository interface {
	AddUserToCompany(userAuth domain.User, user *domain.User, company *domain.Company, roles domain.Roles) error
	SyncUserRolesInCompany(user domain.User, company *domain.Company, roles domain.Roles) error
}

type companyUserRepository struct {
	db *gorm.DB
}

func NewCompanyUserRepository(db *gorm.DB) CompanyUserRepository {
	return &companyUserRepository{db: db}
}

func (r *companyUserRepository) AddUserToCompany(userAuth domain.User, user *domain.User, company *domain.Company, roles domain.Roles) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var existingCompanyUser domain.CompanyUser

		tx.Where("user_id = ? AND company_id = ?", user.ID.String(), company.ID.String()).First(&existingCompanyUser)

		if len(existingCompanyUser.CompanyId) > 0 {
			return errors.New(locales.UserAlreadyAssociatedCompany)
		}

		companyUser := domain.CompanyUser{
			CreatorId: userAuth.ID.String(),
			UserId:    user.ID.String(),
			CompanyId: company.ID.String(),
			Roles:     &roles,
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

func (r companyUserRepository) SyncUserRolesInCompany(user domain.User, company *domain.Company, roles domain.Roles) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var companyUser domain.CompanyUser

		if err := tx.Where("user_id = ? AND company_id = ?", user.ID.String(), company.ID.String()).First(&companyUser).Error; err != nil {
			return err
		}

		var existingRoles []domain.CompanyUserRole
		if err := tx.Where("company_user_id = ?", companyUser.ID).Find(&existingRoles).Error; err != nil {
			return err
		}

		if err := syncRoles(tx, companyUser, roles, existingRoles); err != nil {
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
