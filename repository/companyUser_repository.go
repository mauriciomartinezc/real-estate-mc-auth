package repository

import (
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
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

func (r companyUserRepository) AddUserToCompany(userAuth domain.User, user *domain.User, company *domain.Company, roles domain.Roles) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
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
	var companyUserRoles domain.CompanyUserRoles

	for _, role := range roles {
		companyUserRoles = append(companyUserRoles, domain.CompanyUserRole{
			CompanyUserId: companyUser.ID.String(),
			RoleId:        role.ID.String(),
			CompanyUser:   companyUser,
			Role:          role,
		})
	}

	if err := tx.Create(&companyUserRoles).Error; err != nil {
		return err
	}

	return nil
}

func syncRoles(tx *gorm.DB, companyUser domain.CompanyUser, newRoles domain.Roles, existingRoles domain.CompanyUserRoles) error {
	// Crear mapas para roles actuales y nuevos
	existingRoleMap := make(map[string]bool)
	for _, role := range existingRoles {
		existingRoleMap[role.RoleId] = true
	}

	newRoleMap := make(map[string]bool)
	for _, role := range newRoles {
		newRoleMap[role.ID.String()] = true
		// Agregar nuevos roles si no existen
		if !existingRoleMap[role.ID.String()] {
			if err := tx.Create(&domain.CompanyUserRole{
				CompanyUserId: companyUser.ID.String(),
				RoleId:        role.ID.String(),
			}).Error; err != nil {
				return err
			}
		}
	}

	// Eliminar roles que ya no están en la solicitud
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
