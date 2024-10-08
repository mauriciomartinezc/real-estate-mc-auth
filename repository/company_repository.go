package repository

import (
	"github.com/google/uuid"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"gorm.io/gorm"
)

type CompanyRepository interface {
	Create(company *domain.Company, user domain.User) error
	FindByID(id uuid.UUID) (*domain.Company, error)
	Update(company *domain.Company) error
	AssociateUserToCompany(company *domain.Company, user domain.User, role domain.Role, userCreator domain.User) error
	CompaniesMe(user domain.User) (domain.Companies, error)
}

type companyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) CompanyRepository {
	return &companyRepository{db: db}
}

func (r *companyRepository) Create(company *domain.Company, user domain.User) error {
	if err := r.db.Create(company).Error; err != nil {
		return err
	}

	roleAdmin := new(domain.Role)
	roleAdmin.ID = domain.ROLES["ADMIN"].ID
	userCreator := user

	err := r.AssociateUserToCompany(company, user, *roleAdmin, userCreator)
	if err != nil {
		return err
	}

	return nil
}

func (r *companyRepository) FindByID(id uuid.UUID) (*domain.Company, error) {
	var company domain.Company
	if err := r.db.First(&company, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &company, nil
}

func (r *companyRepository) Update(company *domain.Company) error {
	if err := r.db.Save(company).Error; err != nil {
		return err
	}
	return nil
}

func (r *companyRepository) AssociateUserToCompany(company *domain.Company, user domain.User, role domain.Role, userCreator domain.User) error {
	roles := &domain.Roles{role}

	companyUser := domain.CompanyUser{
		CompanyId: company.ID.String(),
		UserId:    user.ID.String(),
		CreatorId: userCreator.ID.String(),
		Roles:     roles,
	}

	if err := r.db.Create(&companyUser).Error; err != nil {
		return err
	}

	return nil
}

func (r *companyRepository) CompaniesMe(user domain.User) (domain.Companies, error) {
	var companies domain.Companies

	err := r.db.Joins("JOIN company_users ON company_users.company_id = companies.id").
		Where("company_users.user_id = ?", user.ID).
		Find(&companies).Error

	if err != nil {
		return nil, err
	}

	return companies, nil
}
